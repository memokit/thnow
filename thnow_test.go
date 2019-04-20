package thnow

import (
	"log"
	"testing"
	"time"
)

func TestDate(t *testing.T) {

	// location, _ := time.LoadLocation("Asia/Bangkok")

	// dstDay := time.Date(2018, 12, 14, 0, 0, 0, 0, location)

	date := String("14-ธ.ค.-1987").ToDate(1, 2, 3, 0)

	// if xx != dstDay {
	// 	t.Errorf("Failed")
	// }

	log.Println("######### Now Date Time #########")
	log.Println(time.Now())

	log.Println("######### Date Time #########")
	log.Println(date)

	log.Println("######### Set Format Date Time #########")
	log.Println(date.Format("02-Jan-2006 15:04:05 "))
	log.Println(date.Format("02/Jan/2006"))

	dateFormat := Date(date)
	log.Println("######### To String Format Date Time #########")
	dateStr1 := dateFormat.ToString("TH", "F")
	dateStr2 := dateFormat.ToString("EN", "M")
	dateStr3 := dateFormat.ToString("EN", "2006 02 Jan")
	dateStr4 := dateFormat.ToString()
	log.Println(dateStr1)
	log.Println(dateStr2)
	log.Println(dateStr3)
	log.Println(dateStr4)

	log.Println("######### Get Minute #########")
	log.Println(dateFormat.Minute())

	log.Println("######### Get Hour #########")
	log.Println(dateFormat.Hour())

	log.Println("######### Get Day #########")
	log.Println(dateFormat.Day())

	log.Println("######### Get Month #########")
	log.Println(dateFormat.Month())

	log.Println("######### Get Year #########")
	log.Println(dateFormat.Year())
}
