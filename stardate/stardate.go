// stardate.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	// go 1.1.x stdlib only below here
	"flag"
	"fmt"
	"time"
	//
	"github.com/hotei/mdr"
)

var (
	doClock   bool
	doReverse bool
	mo        int
	da        int
	hr        int
	min       int
	month     time.Month
)

func init() {
	flag.BoolVar(&doClock, "clock", false, "clock mode")
	flag.BoolVar(&doReverse, "reverse", false, "clock mode")
	flag.IntVar(&mo, "month", 1, "month")
	flag.IntVar(&da, "day", 1, "day")
	flag.IntVar(&hr, "hr", 0, "hour")
	flag.IntVar(&min, "min", 0, "min")
}

func main() {
	flag.Parse()
	month = time.Month(mo)
	if doReverse {
		now := time.Date(0, month, da, hr, min, 0, 0, time.Local)
		newDate := fmt.Sprintf("%9.4f jday(%d)", mdr.StarDate(now), now.YearDay())
		fmt.Printf("\n\n\n\n\nStardate %s\n", newDate)
	}
	if doClock {
		oldDate := ""
		for {
			now := time.Now()
			newDate := fmt.Sprintf("%9.4f", mdr.StarDate(now))
			if oldDate != newDate {
				fmt.Printf("\n\n\n\n\nStardate %s %s jday(%d)\n", newDate, now.String(), now.YearDay())
				oldDate = newDate
			}
			time.Sleep(60 * time.Second)
		}
	}
	now := time.Now()
	nowDate := fmt.Sprintf("%9.4f", mdr.StarDate(now))
	fmt.Printf("right-now is Stardate(%s) aka %s jday(%d)\n", nowDate, now.String(), now.YearDay())
}
