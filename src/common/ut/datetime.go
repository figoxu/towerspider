package ut

import (
	"github.com/quexer/utee"
	"strings"
	"time"
)

const TOWER_TIME_FMT = "2006-01-02T15:04:05"

func ClearTowerZoneInfo(text string) string { //tower时间范本: 2019-03-25T21:45:41+08:00
	return strings.Replace(text, "+08:00", "", -1)
}

func FormatTime(format, text string) time.Time {
	loc, err := time.LoadLocation("Local")
	utee.Chk(err)
	t, err := time.ParseInLocation(format, text, loc)
	utee.Chk(err)
	return t
}
