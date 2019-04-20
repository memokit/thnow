package thnow

import (
	"strconv"
	"strings"
	"time"
)

/**
* Convet time.Time To Date String format
**/
func (date DateNow) toString(chronology string, pattern string, format string) (string, error) {
	var defaultFormat = "02 Jan 2006 15:04:05"
	var result = time.Now().Format(defaultFormat)
	var err error
	day := date.Day()
	weekday := int(date.Weekday())
	month := int(date.Month())
	year := date.Year()

	if chronology == "TH" {
		if year < 2300 {
			year = year + 543
		}
	} else {
		if year > 2300 {
			year = year - 543
		}
	}

	if format == "M" {
		if chronology == "TH" {
			monthArr := [12]string{"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."}

			result = strconv.Itoa(day) + " " + monthArr[month-1] + " " + strconv.Itoa(year)
		} else {
			monthArr := [12]string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

			result = strconv.Itoa(day) + " " + monthArr[month-1] + " " + strconv.Itoa(year)
		}
	} else if format == "F" {

		if chronology == "TH" {
			dayArr := [7]string{"อาทิตย์", "จันทร์", "อังคาร", "พุธ", "พฤหัสบดี", "ศุกร์", "เสาร์"}
			monthArr := [12]string{
				"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
				"กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"}

			result = "วัน" + dayArr[weekday] + "ที่ " + strconv.Itoa(day) + " " + monthArr[month-1] + " พ.ศ. " + strconv.Itoa(year)
		} else {
			dayArr := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
			monthArr := [12]string{
				"January", "February", "March", "April", "May", "June",
				"July", "August", "September", "October", "November", "December"}

			result = strconv.Itoa(day) + " " + monthArr[month-1] + " " + strconv.Itoa(year) + " " + dayArr[weekday]
		}
	} else {
		if pattern != "" {
			result = date.Format(pattern)
		} else {
			result = date.Format(defaultFormat)
		}
	}

	return result, err
}

/**
* Convet Date String To time.Time
**/
// func toDate(dateStr string, hour int, minute int, second int, millisecond int) (time.Time, error) {
// 	var result = time.Now()
// 	var err error
// 	var day int
// 	var month int
// 	var year int
// 	var dateArr []string

// 	if len(strings.TrimSpace(dateStr)) == 0 {
// 		dateStr = time.Now().Format("02/01/2006")
// 	}

// 	if dateArr = strings.Split(dateStr, "/"); len(dateArr) != 3 {
// 		if dateArr = strings.Split(dateStr, " "); len(dateArr) != 3 {
// 			dateArr = strings.Split(dateStr, "-")
// 		}
// 	}

// 	day, _ = strconv.Atoi(dateArr[0])
// 	month, _ = strconv.Atoi(dateArr[1])
// 	year, _ = strconv.Atoi(dateArr[2])

// 	if len(dateArr) == 3 {

// 		thaiMonthArr := [24]string{
// 			"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค.",
// 			"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

// 		for index, m := range thaiMonthArr {
// 			if dateArr[1] == m {
// 				month = index + 1

// 				if month > 12 {
// 					month = month - 12
// 				}
// 				break
// 			}
// 		}
// 	}

// 	if year > 2300 {
// 		year = year - 543
// 	}

// 	result = time.Date(year, time.Month(month), day, hour, minute, second, millisecond, time.UTC)

// 	return result, err

// }

// toD Convet Date String To time.Time
// dateStr
func (dateStr StringNow) toDate(optional ...int) (time.Time, error) {
	var result = time.Now()
	var err error
	var day, month, year int
	var dateArr []string
	hour, minute, second, millisecond := 0, 0, 0, 0

	if len(dateStr.string) == 0 {
		dateStr.string = time.Now().Format("02/01/2006")
	}

	if dateArr = strings.Split(dateStr.string, "/"); len(dateArr) != 3 {
		if dateArr = strings.Split(dateStr.string, " "); len(dateArr) != 3 {
			dateArr = strings.Split(dateStr.string, "-")
		}
	}

	day, _ = strconv.Atoi(dateArr[0])
	month, _ = strconv.Atoi(dateArr[1])
	year, _ = strconv.Atoi(dateArr[2])

	if len(dateArr) == 3 {

		thaiMonthArr := [24]string{
			"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค.",
			"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

		for index, m := range thaiMonthArr {
			if dateArr[1] == m {
				month = index + 1

				if month > 12 {
					month = month - 12
				}
				break
			}
		}
	}

	if year > 2300 {
		year = year - 543
	}

	if len(optional) > 0 && len(optional) <= 1 {
		hour = optional[0]
	} else if len(optional) > 1 && len(optional) <= 2 {
		hour = optional[0]
		minute = optional[1]
	} else if len(optional) > 2 && len(optional) <= 3 {
		hour = optional[0]
		minute = optional[1]
		second = optional[2]
	} else if len(optional) > 3 && len(optional) <= 4 {
		hour = optional[0]
		minute = optional[1]
		second = optional[2]
		millisecond = optional[3]
	}

	result = time.Date(year, time.Month(month), day, hour, minute, second, millisecond, time.UTC)

	return result, err
}
