// SimpleSNTP_test.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
	"mdr"
	"net"
	"os"
	"testing"
	"time"
)

// ask tock (USNO time server) for time, reference http://www.usno.navy.mil/USNO/time/ntp/eastern-tz
// use [tick | tock | ntp2 ] as source for NTP V4 info
//or if you can, use one closer to you.

func Test_SNTP(t *testing.T) {
	fmt.Printf("<start SimpleSNMP_test.go>\n")

	fmt.Printf("Time 1970 = %g years\n", TIME1970/sec_in_yr)

	udpaddr, err := net.ResolveUDPAddr("udp4", "tock.usno.navy.mil:123")
	if err != nil {
		fmt.Printf("UDP parse error %v\n", err)
		os.Exit(1)
	}

	con, err := net.DialUDP("udp4", nil, udpaddr)
	if err != nil {
		fmt.Printf("UDP connect error %v\n", err)
		os.Exit(1)
	}

	timeReq := make([]byte, 48)
	timeReq[0] = 0x1b
	n, err := con.Write(timeReq)
	if err != nil {
		fmt.Printf("UDP write error %v\n", err)
		os.Exit(1)
	}
	if n != 48 {
		fmt.Printf("Not enough bytes written to UDP\n")
		os.Exit(1)
	}
	data := make([]byte, 1024)
	n, from_addr, err := con.ReadFromUDP(data)
	if err != nil {
		fmt.Printf("UDP read error %v\n", err)
		os.Exit(1)
	}
	if n != 48 {
		fmt.Printf("Not enough bytes read from UDP\n")
		os.Exit(1)
	}

	fmt.Printf("Read %d bytes from %v\n", n, from_addr)
	//	fmt.Printf("Data: %v\n", data)
	nowTime := time.Now()
	now := nowTime.Unix()
	ntpSec := mdr.ThirtyTwoNet(data[40:44]) // convert BIGendian value to int32
	fmt.Printf("NTP Secs since 1970 Jan1 = %16x\n", ntpSec)
	fmt.Printf("Now                      = %16x\n", now)
	fmt.Printf("Time of 1970-01-01       = %16x\n", int64(TIME1970))
	fmt.Printf("Now + 1970               = %16x\n", now+TIME1970)

	localTime := time.Unix(int64(ntpSec)-TIME1970, 0)
	fmt.Printf("Go says local time   is                            : %s\n", nowTime.String())
	fmt.Printf("SNTP from tock.usno.navy.mil:123 says local time is: %s\n", localTime.String())

	diff := nowTime.Sub(localTime)
	diffsec := diff.Seconds()
	if diffsec < 0 {
		diff = -diff
	}
	fmt.Printf("diff = %v\n", diffsec)
	if diffsec > 10 {
		t.Errorf("time reported by go doesn't match SNTP from tock.usno.navy.mil")
	}
	fmt.Printf("<End SimpleSNMP.go>\n")

}
