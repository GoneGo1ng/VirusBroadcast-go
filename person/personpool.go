package person

import (
	"virusbroadcast/constants"
	"virusbroadcast/point"
	"virusbroadcast/util"
)

// 人员池
type Pool struct {
	Persons []*Person
}

// 获取指定状态的人数
func (pp *Pool) GetPeopleSize(state int) int {
	if state == -1 {
		return len(pp.Persons)
	}
	i := 0
	for _, person := range pp.Persons {
		if person.State == state {
			i++
		}
	}
	return i
}

var pool *Pool

// 人员池实例
func GetInstance() *Pool {
	if pool != nil {
		return pool
	}

	pool = &Pool{Persons: []*Person{}}

	// 根据标准正态分布生成随机人口初始位置和目标位置
	for i := 0; i < constants.CityPersonSize; i++ {
		x := int(util.StdGaussian(constants.InitialPointSig, constants.CityWidth/2))
		y := int(util.StdGaussian(constants.InitialPointSig, constants.CityHeight/2))
		if x > constants.CityWidth {
			x = constants.CityWidth
		}
		if y > constants.CityHeight {
			y = constants.CityHeight
		}
		// 初始化人员，将人员信息放入人员池中
		pool.Persons = append(pool.Persons, &Person{
			TargetXU: util.StdGaussian(constants.TargetUSig, constants.CityWidth/2),
			TargetYU: util.StdGaussian(constants.TargetUSig, constants.CityHeight/2),
			Point: &point.Point{
				X: x,
				Y: y,
			},
		})
	}

	return pool
}
