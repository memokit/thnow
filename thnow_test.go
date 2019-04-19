package thnow

import (
	"log"
	"testing"
	"time"
)

func TestDate(t *testing.T) {

	// location, _ := time.LoadLocation("Asia/Bangkok")

	// dstDay := time.Date(2018, 12, 14, 0, 0, 0, 0, location)

	xx, _ := toDate("14-ธ.ค.-1987", 15, 0, 0, 0)

	// if xx != dstDay {
	// 	t.Errorf("Failed")
	// }

	log.Println("+++++++++++++++++++++")
	log.Println(time.Now())

	log.Println("---------------------")
	log.Println(xx)
	log.Println(xx.Format("02-Jan-2006 15:04:05 "))
	log.Println(xx.Format("02-Jan-2006"))

	log.Println("##################")

	tdStr, _ := toString(xx, "EN", "02 Jan 2006", "F")
	log.Println(tdStr)

}
