package hospital

import (
	"virusbroadcast/bed"
	"virusbroadcast/constants"
	"virusbroadcast/point"
)

// 医院
type Hospital struct {
	Point *point.Point // 坐标
	Beds  []*bed.Bed   // 床位
}

var hospital *Hospital

// 医院实例
func GetInstance() *Hospital {
	if hospital != nil {
		return hospital
	}

	// 初始化医院
	hospital = &Hospital{
		Point: &point.Point{
			X: constants.HospitalX,
			Y: constants.HospitalY,
		},
		Beds: nil,
	}

	// 有几列床位（一列五十行）
	column := constants.BedCount / 50

	// 初始化床位
	for i := 0; i < column; i++ {
		for j := 0; j < 50; j++ {
			bed := &bed.Bed{
				Point: &point.Point{
					X: hospital.Point.X + i*6,
					Y: hospital.Point.Y + j*6,
				},
				IsEmpty: true,
			}
			hospital.Beds = append(hospital.Beds, bed)
		}
	}

	return hospital
}

// 使用一个空床位
func (h *Hospital) PickBed() *bed.Bed {
	for _, bed := range h.Beds {
		if bed.IsEmpty {
			return bed
		}
	}
	return nil
}
