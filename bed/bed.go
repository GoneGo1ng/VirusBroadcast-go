package bed

import "virusbroadcast/point"

// 床位
type Bed struct {
	Point   *point.Point // 图上的位置
	IsEmpty bool         // 是否占用了该床位
}

// 归还床位
func (bed *Bed) ReturnBed() {
	if bed != nil {
		bed.IsEmpty = true
	}
}
