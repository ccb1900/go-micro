package timex

import "time"

// 固定时间在东八区
func Fixed() {
	time.Local = time.FixedZone("CST", 8*3600)
}
