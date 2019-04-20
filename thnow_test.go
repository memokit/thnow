package thnow

import (
	"log"
	"testing"
	"time"
)

func TestDate(t *testing.T) {

	// location, _ := time.LoadLocation("Asia/Bangkok")

	// dstDay := time.Date(2018, 12, 14, 0, 0, 0, 0, location)

	xx, _ := String("14-ธ.ค.-1987").ToDate(15, 0, 0, 0)

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

	tdStr, _ := Date(xx).ToString("EN", "F", "02 Jan 2006")
	tdStr2, _ := Date(xx).ToString("EN", "F")
	tdStr3, _ := Date(xx).ToString("EN", "2006 02 Jan")
	tdStr4, _ := Date(xx).ToString()
	log.Println(tdStr)
	log.Println(tdStr2)
	log.Println(tdStr3)
	log.Println(tdStr4)

	log.Println("**********************")

	yy, _ := String("14-ธ.ค.-2560").ToDate()
	log.Println(yy.Format("02-Jan-2006 15:04:05 "))

}
