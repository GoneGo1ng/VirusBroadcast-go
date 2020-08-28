package person

import "virusbroadcast/point"

// 移位目标
type Target struct {
	Point   *point.Point // 在图上的位置
	Arrived bool         // 是否到达目标
}
