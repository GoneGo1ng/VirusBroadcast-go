package constants

const (
	NORMAL    = iota // 正常人，未感染的健康人
	CURED            // 治愈的
	SHADOW           // 潜伏期
	CONFIRMED        // 发病且已确诊为感染病人
	FREEZE           // 隔离治疗，禁止位移
	DEATH            // 病死者
)

const OriginalCount = 50       // 初始感染数量
const BroadRate = 0.8          // 传播率
const ShadowSig = 4            // 潜伏期方差sigma
const ShadowTime = 14          // 潜伏时间均值u
const SafeDist = 2             // 感染距离
const HospitalReceiveTime = 10 // 医院收治响应时间
const BedCount = 500           // 医院床位

// 流动意向平均值，建议调整范围：[-0.99,0.99]
// -0.99 人群流动最慢速率，甚至完全控制疫情传播
// 0.99为人群流动最快速率, 可导致全城感染
const U = 0.99
const TargetSig = 50        // 目标位置方差sigma
const TargetUSig = 150      // 目标位置均值方差sigma
const InitialPointSig = 100 // 初始位置方差sigma
const DeathPointSig = 40    // 死亡人员位置方差sigma

const CityPersonSize = 5000 // 城市总人口数量
const FatalityRate = 0.50   // fatality_rate病死率（病死数/确诊数）
const DieTime = 30          // 死亡时间均值，从发病（确诊）时开始计时
const DieVariance = 5       // 死亡时间方差
const CureTime = 50         // 治愈时间均值，从发病（确诊）时开始计时
const CureVariance = 5      // 治愈时间方差

// 城市大小即窗口边界，限制不允许出城
const CityWidth = 800
const CityHeight = 800

// 医院起始点坐标
const HospitalX = 840
const HospitalY = 140
