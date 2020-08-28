package util

import (
	"math/rand"
)

// 标准正态分布化
// 流动意愿标准化后判断是在0的左边还是右边从而决定是否流动。
// 设X随机变量为服从正态分布，sigma是影响分布形态的系数 u值决定正态分布均值
// 推导：
// StdX = (X-u)/sigma
// X = sigma * StdX + u
// sigma 正态标准差sigma值
// u     正态均值参数mu
func StdGaussian(sigma, u float64) float64 {
	// rand.Seed(time.Now().UnixNano())
	x := rand.NormFloat64()
	return sigma*x + u
}
