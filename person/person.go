package person

import (
	"math"
	"math/rand"
	"virusbroadcast/bed"
	"virusbroadcast/constants"
	"virusbroadcast/global"
	"virusbroadcast/hospital"
	"virusbroadcast/point"
	"virusbroadcast/util"
)

// 会随机移动的人员
type Person struct {
	Point    *point.Point // 在图上的位置
	Target   *Target      // 移动目标位置
	TargetXU float64      // x轴的均值mu
	TargetYU float64      // y轴的均值mu
	State    int          // 状态

	InfectedTime  int // 感染时刻
	ConfirmedTime int // 确诊时刻
	CureMoment    int // 治愈时刻
	DieMoment     int // 死亡时刻，为0代表未确定，-1代表会治愈

	UseBed *bed.Bed // 床位使用情况
}

// 流动意愿标准化
// 根据标准正态分布生成随机人口流动意愿
// 流动意愿标准化后判断是在0的左边还是右边从而决定是否流动
// 设X随机变量为服从正态分布，sigma是影响分布形态的系数，从而影响整体人群流动意愿分布
// u值决定正态分布的中轴是让更多人群偏向希望流动或者希望懒惰。
func (p *Person) WantMove() bool {
	return util.StdGaussian(1, constants.U) > 0
}

// 是否感染
func (p *Person) IsInfected() bool {
	return p.State >= constants.SHADOW
}

// 被感染
func (p *Person) BeInfected() {
	p.State = constants.SHADOW
	p.InfectedTime = global.WorldTime
}

// 计算两点之间的直线距离
func (p *Person) Distance(op *Person) float64 {
	return math.Sqrt(math.Pow(float64(p.Point.X-op.Point.X), 2) +
		math.Pow(float64(p.Point.Y-op.Point.Y), 2))
}

// 不同状态下的单个人实例运动行为
func (p *Person) Action() {
	// 如果处于隔离或者死亡状态，则无法行动
	if p.State == constants.FREEZE || p.State == constants.DEATH {
		return
	}
	// 如果无移动意愿，则无法运动
	if !p.WantMove() {
		return
	}

	// 存在流动意愿的，将进行流动，流动位移仍然遵循标准正态分布
	if p.Target == nil || p.Target.Arrived {
		targetX := util.StdGaussian(constants.TargetSig, p.TargetXU)
		targetY := util.StdGaussian(constants.TargetSig, p.TargetYU)
		p.Target = &Target{
			Point: &point.Point{
				X: int(targetX),
				Y: int(targetY),
			},
			Arrived: false,
		}
	}

	// 计算移动距离
	dx := p.Target.Point.X - p.Point.X
	dy := p.Target.Point.Y - p.Point.Y
	length := math.Sqrt(math.Pow(float64(dx), 2) + math.Pow(float64(dy), 2))
	// 判断是否到达目标位置
	if length < 1 {
		p.Target.Arrived = true
		return
	}

	// x轴dX为位移量，符号为沿x轴前进方向, 即udX为X方向表示量
	udx := int(float64(dx) / length)
	if udx == 0 && dx != 0 {
		if dx > 0 {
			udx = 1
		} else {
			udx = -1
		}
	}

	// y轴dY为位移量，符号为沿x轴前进方向，即udY为Y方向表示量
	udy := int(float64(dy) / length)
	if udy == 0 && dy != 0 {
		if dy > 0 {
			udy = 1
		} else {
			udy = -1
		}
	}

	// 如果超出了边界，说明移动目标在边界外，将移动目标置空并往放方向移动
	if p.Point.X > constants.CityWidth || p.Point.X < 0 {
		p.Target = nil
		if udx > 0 {
			udx = -udx
		}
	}

	// 如果超出了边界，说明移动目标在边界外，将移动目标置空并往放方向移动
	if p.Point.Y > constants.CityHeight || p.Point.Y < 0 {
		p.Target = nil
		if udy > 0 {
			udy = -udy
		}
	}

	p.Point.MoveTo(udx, udy)
}

// 对各种状态的人进行不同的处理，更新发布市民健康状态
func (p *Person) Update() {
	// rand.Seed(time.Now().UnixNano())
	// 如果已死亡，则不处理
	if p.State == constants.DEATH {
		return
	}
	// 处理治愈的感染者
	// 如果到了治愈时刻，则将人员状态改为治愈，移动到人群中，并归还床位
	if (p.State == constants.CONFIRMED || p.State == constants.FREEZE) &&
		global.WorldTime >= p.CureMoment && p.CureMoment > 0 {
		p.State = constants.CURED
		p.Point.X = int(util.StdGaussian(constants.InitialPointSig, float64(constants.CityWidth/2)))
		p.Point.Y = int(util.StdGaussian(constants.InitialPointSig, float64(constants.CityHeight/2)))
		p.UseBed.ReturnBed()
	}
	// 处理死亡的感染者
	// 如果到了死亡时刻，则将人员状态改为死亡，并归还床位
	if (p.State == constants.CONFIRMED || p.State == constants.FREEZE) &&
		global.WorldTime >= p.DieMoment && p.DieMoment > 0 {
		// 如果死亡时在医院，则移动到人群中（为了跟直观到展现死亡人数）
		if p.State == constants.FREEZE {
			p.Point.X = int(util.StdGaussian(constants.DeathPointSig, float64(constants.CityWidth/2)))
			p.Point.Y = int(util.StdGaussian(constants.DeathPointSig, float64(constants.CityWidth/2)))
		}
		p.State = constants.DEATH
		p.UseBed.ReturnBed()
	}
	// 处理已经确诊的感染者
	if p.State == constants.CONFIRMED && p.DieMoment == 0 {
		// 随机一个数字，如果数字小于死亡率，则设置该人员的死亡时间，否则设置该人员的治愈时间
		if rand.Float64() <= constants.FatalityRate {
			// 希望通过国家、政府、白衣天使以及所有爱心人生的努力，永远都不要进这个判断
			dieTime := int(util.StdGaussian(constants.DieVariance, constants.DieTime))
			p.DieMoment = p.ConfirmedTime + dieTime
		} else {
			cureTime := int(util.StdGaussian(constants.CureVariance, constants.CureTime))
			p.CureMoment = p.ConfirmedTime + cureTime
			p.DieMoment = -1
		}
	}
	// 如果患者已经确诊，且（世界时刻-确诊时刻）大于医院响应时间，即医院准备好病床了，可以抬走了
	if p.State == constants.CONFIRMED &&
		global.WorldTime-p.ConfirmedTime >= constants.HospitalReceiveTime {
		hospital := hospital.GetInstance()
		bed := hospital.PickBed() // 查找空床位
		if bed == nil {

		} else {
			// 住院
			p.UseBed = bed
			p.State = constants.FREEZE
			p.Point.X = bed.Point.X
			p.Point.Y = bed.Point.Y
			bed.IsEmpty = false
		}
	}
	// 增加一个正态分布用于潜伏期内随机发病时间
	stdRnShadowTime := int(util.StdGaussian(constants.ShadowSig, constants.ShadowTime))
	if global.WorldTime-p.InfectedTime > stdRnShadowTime && p.State == constants.SHADOW {
		p.State = constants.CONFIRMED
		p.ConfirmedTime = global.WorldTime
	}
	// 处理未隔离者的移动问题
	p.Action()
	// 非正常人不会被感染
	if p.State >= constants.CURED {
		return
	}
	// 通过一个随机值和安全距离决定感染其他人（正常人）
	for _, np := range GetInstance().Persons {
		if (np.State == constants.SHADOW || np.State == constants.CONFIRMED) &&
			rand.Float64() < constants.BroadRate && p.Distance(np) < constants.SafeDist {
			p.State = constants.SHADOW
			p.InfectedTime = global.WorldTime
			break
		}
	}
}
