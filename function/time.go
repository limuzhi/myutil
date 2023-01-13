/*
 * @PackageName: utils
 * @Description:时间
 * @Author: limuzhi
 * @Date: 2022/12/1 11:33
 */

package function

import "time"

const (
	CSTLayout        = "2006-01-02 15:04:05"
	TimeFormatDateV1 = "2006-01-02"
)

//当前时间的时间戳
func GetCurrentTime() int64 {
	return time.Now().Unix()
}

//今天零点 时间戳秒
func GetTodayZeroUnix() int64 {
	ct := time.Now()
	return time.Date(ct.Year(), ct.Month(), ct.Day(), 0, 0, 0, 0, ct.Location()).Unix()
}

//日期转时间Time
func FormatDateToTime(date string) time.Time {
	loc, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation(CSTLayout, date, loc)
	return t
}

//日期转时间戳
func FormDateToUnix(date string) int64 {
	t := FormatDateToTime(date)
	return t.Unix()
}

//时间Time转字符串
func TimeToFormDate(t time.Time) string {
	d := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(),
		t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	return d.Format(CSTLayout)
}

//时间Time转时间戳
func TimeToUnix(t time.Time) int64 {
	return t.Unix()
}

//时间戳转日期
func UnixToDate(t int64) string {
	return time.Unix(t, 0).Format(CSTLayout)
}

//时间戳转日期
func UnixToTime(t int64) time.Time {
	return time.Unix(t, 0)
}

//本月1号开始时间
func StartOfThisMonth() time.Time {
	ct := time.Now()
	return time.Date(ct.Year(), ct.Month(), 1, 0, 0, 0, 0, ct.Location())
}

//下月1号
func StartOfLastMonth() time.Time {
	ct := time.Now()
	return time.Date(ct.Year(), ct.Month()+1, 1, 0, 0, 0, 0, ct.Location())
}

//获取当前天距离明天凌晨时间差
func GetTodaySurplusSecond() int64 {
	t := time.Now()
	timeStr := t.Format(TimeFormatDateV1)
	t1, _ := time.ParseInLocation(TimeFormatDateV1, timeStr, time.Local)
	d1 := t1.AddDate(0, 0, 1)
	dff := d1.Sub(t).Seconds()
	return int64(dff)
}
