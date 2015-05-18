// sntp.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

// NTP uses an epoch of January 1, 1900
//
// The 64-bit timestamps used by NTP consist of a 32-bit seconds part
//  	and a 32-bit fractional second part  (fraction is unused here)
// The magic number slice [40:44] picks out the 32bit seconds part of the msg
// 		which is then converted from Net Order (BigEndian) to Intel (LittleEndian)
//
// A better approach might be to inspect ENV array and convert or not
// 		based on value of GOARCH   (amd64 | ? )
//
//	pick timeServerName from hosts listed at
//		//support.ntp.org/bin/view/Servers/StratumTwoTimeServers
// 		Check http://en.wikipedia.org/wiki/Network_Time_Protocol for history
// 		http://www.usno.navy.mil/USNO/time/ntp lists hosts available for .mil
// 			such as "tock.usno.navy.mil:123"
package main

// NOTE: This is tested to work only on LittleEndian (Intel/AMD) CPUs

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	// 1970 Jan 1, 0000 is the UNIX epoch
	TIME1970  = 2208988800 // seconds since Jan 1, 1900  (NTP Epoch)
	sec_in_yr = (86400 * 365.25)
)

var (
	timeServerName      = "ntp-2.vt.edu" + ":123"
	altServerName       = "ntp.theforest.us" + ":123"
	verbose        bool = false
)

// convert from BIG endian four byte slice to int32
func UThirtyTwoFromNet(n []byte) uint32 {
	if len(n) != 4 {
		fmt.Printf("Slice must be exactly 4 bytes\n")
		os.Exit(1)
	}
	var rc uint32
	rc = uint32(n[0])
	rc <<= 8
	rc |= uint32(n[1])
	rc <<= 8
	rc |= uint32(n[2])
	rc <<= 8
	rc |= uint32(n[3])
	return rc
}

func main() {
	if verbose {
		fmt.Printf("<start sntp.go>\n")
		fmt.Printf("Time 1970 = %g years\n", TIME1970/sec_in_yr)
	}
	udpaddr, err := net.ResolveUDPAddr("udp4", altServerName)
	if err != nil {
		fmt.Printf("UDP parse error %v\n", err)
		os.Exit(1)
	}
	con, err := net.DialUDP("udp4", nil, udpaddr)
	if err != nil {
		fmt.Printf("UDP connect error %v\n", err)
		os.Exit(2)
	}
	timeReq := make([]byte, 48)
	timeReq[0] = 0x1b
	n, err := con.Write(timeReq)
	if err != nil {
		fmt.Printf("UDP write error %v\n", err)
		os.Exit(3)
	}
	if n != 48 {
		fmt.Printf("Not enough bytes written to UDP\n")
		os.Exit(4)
	}
	data := make([]byte, 1024)
	n, from_addr, err := con.ReadFromUDP(data)
	if err != nil {
		fmt.Printf("UDP read error %v\n", err)
		os.Exit(-1)
	}
	if n != 48 {
		fmt.Printf("Not enough bytes read from UDP\n")
		os.Exit(-2)
	}
	fmt.Printf("Read %d bytes from %v\n", n, from_addr)
	if verbose {
		fmt.Printf("Data: %v\n", data)
	}
	nowTime := time.Now()
	now := nowTime.Unix()
	ntpSec := UThirtyTwoFromNet(data[40:44]) // convert BIGendian value to int32
	if verbose {
		fmt.Printf("NTP Secs since 1970 Jan1 = %16x\n", ntpSec)
		fmt.Printf("Now                      = %16x\n", now)
		fmt.Printf("Time of 1970-01-01       = %16x\n", int64(TIME1970))
		fmt.Printf("Now + 1970               = %16x\n", now+TIME1970)
	}
	localTime := time.Unix(int64(ntpSec)-TIME1970, 0)
	fmt.Printf("Go says local time is : %s\n", nowTime.String())
	fmt.Printf("SNTP from %s says local time is: %s\n", timeServerName, localTime.String())
	diff := nowTime.Sub(localTime)
	diffsec := diff.Seconds()
	if diffsec < 0 {
		diff = -diff
	}
	fmt.Printf("Times differ by %v seconds\n", diffsec)
	if diffsec > 15 {
		fmt.Printf("time reported by go doesn't match SNTP from %s", timeServerName)
		fmt.Printf("May need to check local sntp daemon operation\n")
	}
	if verbose {
		fmt.Printf("<end sntp.go>\n")
	}
}
