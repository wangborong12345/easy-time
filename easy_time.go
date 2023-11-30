package etime

import "time"

// AtStartOfDay 重置时间到 t 的当前天的第一秒
func AtStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// AtEndOfDay 重置时间到 t 的当前天最后一纳秒
func AtEndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}
