package kw_time

import (
	"fmt"
	"strconv"
	"time"
)

// ParseDuration parses a duration string, adding
// support for the "d" unit meaning number of days,
// where a day is assumed to be 24h.
func ParseDuration(s string) (time.Duration, error) {
	var inNumber bool
	var numStart int
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == 'd' {
			daysStr := s[numStart:i]
			days, err := strconv.ParseFloat(daysStr, 64)
			if err != nil {
				return 0, err
			}
			hours := days * 24.0
			hoursStr := strconv.FormatFloat(hours, 'f', -1, 64)
			s = s[:numStart] + hoursStr + "h" + s[i+1:]
			i--
			continue
		}
		if !inNumber {
			numStart = i
		}
		inNumber = (ch >= '0' && ch <= '9') || ch == '.' || ch == '-' || ch == '+'
	}
	return time.ParseDuration(s)
}

var humanizedDurationKey = []string{
	"yy",
	"y",
	"MM",
	"M",
	"dd",
	"d",
	"hh",
	"h",
	"mm",
	"m",
	"ss",
	"s",
}

var humanizedDurationThreshold = map[string]int64{
	"s":  44,           // 0 to 44 seconds
	"ss": 45,           // unset
	"m":  89,           // 45 to 89 seconds
	"mm": 60 * 44,      // 90 seconds to 44 minutes
	"h":  60 * 89,      // 45 to 89 minutes
	"hh": 60 * 60 * 21, // 90 minutes to 21 hours
	"d":  60 * 60 * 35, // 22 to 35 hours
	"dd": 86400 * 25,   // 36 hours to 25 days
	"M":  86400 * 45,   // 26 to 45 days
	"MM": 86400 * 319,  // 45 to 319 days
	"y":  86400 * 547,  // 320 to 547 days (1.5 years)
	"yy": 86400 * 548,  // 548 days+
}

var humanizedDurationUnit = map[string]int64{
	"s":  0,           // 0 to 44 seconds
	"ss": 0,           // unset
	"m":  60,          // 45 to 89 seconds
	"mm": 60,          // 90 seconds to 44 minutes
	"h":  60 * 60,     // 45 to 89 minutes
	"hh": 60 * 60,     // 90 minutes to 21 hours
	"d":  86400,       // 22 to 35 hours
	"dd": 86400,       // 36 hours to 25 days
	"M":  86400 * 30,  // 26 to 45 days
	"MM": 86400 * 30,  // 45 to 319 days
	"y":  86400 * 365, // 320 to 547 days (1.5 years)
	"yy": 86400 * 365, // 548 days+
}

var humanizedDurationText = map[string]string{
	"s":  "几秒前",
	"ss": "%d 秒前",
	"m":  "1 分钟前",
	"mm": "%d 分钟前",
	"h":  "1 小时前",
	"hh": "%d 小时前",
	"d":  "1 天前",
	"dd": "%d 天前",
	"M":  "1 个月前",
	"MM": "%d 个月前",
	"y":  "1 年前",
	"yy": "%d 年前",
}

// HumanizeDurationFrom display time in humanized way
func HumanizeDurationFrom(moment time.Time, now time.Time) string {
	var momentUnix = moment.Unix()
	var nowUnix = now.Unix()

	if momentUnix > nowUnix {
		return moment.Format("2006-01-02")
	}

	var diff = nowUnix - momentUnix
	var key = "yy"
	for _, k := range humanizedDurationKey {
		if diff < humanizedDurationThreshold[k] {
			key = k
		}
	}

	if len(key) == 1 {
		return humanizedDurationText[key]
	}

	var unit = humanizedDurationUnit[key]
	if unit == 0 {
		return fmt.Sprintf(humanizedDurationText[key], diff)
	} else {
		var reminder = diff / unit
		return fmt.Sprintf(humanizedDurationText[key], reminder)
	}
}

// HumanizeDurationFromNow display time in humanized way
func HumanizeDurationFromNow(moment time.Time) string {
	return HumanizeDurationFrom(moment, time.Now())
}

func GetNowDateString() string {
	return time.Now().Format("2006-01-02 15:04:05.999")
}

func DateTimeToString(dateTime time.Time) string {
	return dateTime.Format("2006-01-02 15:04:05")
}

var timeTemplates = []string{
	"2006-01-02 15:04:05", //常规类型
	"2006/01/02 15:04:05",
	"2006-01-02",
	"2006/01/02",
	"15:04:05",
}

// StringToGoTime /* 时间格式字符串转换 */
func StringToGoTime(tm string) time.Time {
	for i := range timeTemplates {
		t, err := time.ParseInLocation(timeTemplates[i], tm, time.Local)
		if nil == err && !t.IsZero() {
			return t
		}
	}
	return time.Time{}
}
