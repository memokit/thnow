package thnow

import (
	"strconv"
	"strings"
	"time"
)

// ToString Convet time.Time To Date String format
func (date DateNow) ToString(optional ...string) string {
	var defaultFormat = "02 Jan 2006 15:04:05"
	var result = time.Now().Format(defaultFormat)
	day := date.Day()
	weekday := int(date.Weekday())
	month := int(date.Month())
	year := date.Year()

	chronology, format, pattern := "TH", "M", ""

	if len(optional) > 0 && len(optional) <= 1 {
		if optional[0] == "EN" || optional[0] == "TH" {
			chronology = optional[0]
		}

	} else if len(optional) > 1 && len(optional) <= 2 {
		chronology = optional[0]

		if optional[1] == "M" || optional[1] == "F" {
			format = optional[1]
		} else {
			pattern = optional[1]
		}

	} else if len(optional) > 2 && len(optional) <= 3 {
		chronology = optional[0]
		format = optional[1]
		pattern = optional[2]
	}

	if chronology == "TH" {
		if year < 2300 {
			year = year + 543
		}
	} else {
		if year > 2300 {
			year = year - 543
		}
	}

	if format == "M" && pattern == "" {
		if chronology == "TH" {
			monthArr := [12]string{"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."}

			result = strconv.Itoa(day) + " " + monthArr[month-1] + " " + strconv.Itoa(year)
		} else {
			monthArr := [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

			result = strconv.Itoa(day) + " " + monthArr[month-1] + " " + strconv.Itoa(year)
		}
	} else if format == "F" && pattern == "" {

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

	return result
}

// ToDate Convet Date String To time.Time
func (dateStr StringNow) ToDate(optional ...int) time.Time {
	var result = time.Now()
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
			"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

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

	return result
}

// Minute get minute
func (date DateNow) Minute() int {
	return date.Time.Minute()
}

// Hour get hour
func (date DateNow) Hour() int {
	return date.Time.Hour()
}

// Day get day
func (date DateNow) Day() int {
	return date.Time.Day()
}

// Month get month
func (date DateNow) Month() int {
	return int(date.Time.Month())
}

// Year get year
func (date DateNow) Year() int {
	return int(date.Time.Year())
}
