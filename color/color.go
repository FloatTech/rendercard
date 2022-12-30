// Package color ...
package color

import "math/rand"

// Randcolor 随机和风颜色
func Randcolor() (r, g, b int) {
	switch rand.Intn(6) {
	case 0: // 红
		r, g, b = rand.Intn(50)+180, rand.Intn(30), rand.Intn(80)+40
	case 1: // 橙
		r, g, b = rand.Intn(40)+210, rand.Intn(50)+70, rand.Intn(50)+20
	case 2: // 黄
		r, g, b = rand.Intn(40)+210, rand.Intn(50)+170, rand.Intn(110)+40
	case 3: // 绿
		r, g, b = rand.Intn(60)+80, rand.Intn(80)+140, rand.Intn(60)+80
	case 4: // 蓝
		r, g, b = rand.Intn(60)+80, rand.Intn(50)+170, rand.Intn(50)+170
	case 5: // 紫
		r, g, b = rand.Intn(60)+80, rand.Intn(60)+60, rand.Intn(50)+170
	}
	return
}
