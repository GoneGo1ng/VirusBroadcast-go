package panel

import (
	"bytes"
	"fmt"
	"github.com/golang/freetype"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math"
	"virusbroadcast/constants"
	"virusbroadcast/global"
	"virusbroadcast/person"
)

// 绘制图片
func Paint() bytes.Buffer {
	// 获取人员实例
	pp := person.GetInstance()

	// 设置背景
	img := image.NewRGBA(image.Rect(0, 0, 1200, 800))
	draw.Draw(img, img.Bounds(), image.NewUniform(color.RGBA{R: 68, G: 68, B: 68, A: 255}), image.ZP, draw.Src)

	// 设置字体
	fontBytes, _ := ioutil.ReadFile("font.otf")
	font, _ := freetype.ParseFont(fontBytes)
	f := freetype.NewContext()
	f.SetDPI(72)
	f.SetFont(font)
	f.SetFontSize(18)
	f.SetClip(img.Bounds())
	f.SetDst(img)

	// 绘制医院
	f.SetSrc(image.NewUniform(color.RGBA{G: 255, A: 255}))
	pt := freetype.Pt(img.Bounds().Dx()-350, img.Bounds().Dy()-700)
	f.DrawString("医院", pt)
	r := image.Rect(constants.HospitalX-20, constants.HospitalY-20, constants.HospitalX+76, constants.HospitalY+316)
	draw.Draw(img, r, image.NewUniform(color.RGBA{R: 105, G: 105, B: 105, A: 255}), image.ZP, draw.Src)

	// 绘制代表人类的圆点
	for _, p := range pp.Persons {
		p.Update()
		c := color.RGBA{R: 221, G: 221, B: 221, A: 255}
		switch p.State {
		case constants.NORMAL:
			c = color.RGBA{R: 221, G: 221, B: 221, A: 255}
		case constants.CURED:
			c = color.RGBA{R: 204, G: 187, B: 204, A: 255}
		case constants.SHADOW:
			c = color.RGBA{R: 255, G: 238, A: 255}
		case constants.CONFIRMED:
			c = color.RGBA{R: 255, A: 255}
		case constants.FREEZE:
			c = color.RGBA{R: 72, G: 255, B: 255, A: 255}
		case constants.DEATH:
			c = color.RGBA{A: 255}
		}
		r := image.Rect(p.Point.X, p.Point.Y, p.Point.X+2, p.Point.Y+2)
		draw.Draw(img, r, image.NewUniform(c), image.ZP, draw.Src)
	}

	// 绘制各种人数及床位数的信息
	f.SetSrc(image.NewUniform(color.White))
	pt = freetype.Pt(img.Bounds().Dx()-250, img.Bounds().Dy()-700)
	f.DrawString(fmt.Sprintf("城市总人数：%d", constants.CityPersonSize), pt)
	f.SetSrc(image.NewUniform(color.RGBA{R: 211, G: 211, B: 211, A: 255}))
	pt = freetype.Pt(img.Bounds().Dx()-250, img.Bounds().Dy()-670)
	f.DrawString(fmt.Sprintf("健康者人数：%d", pp.GetPeopleSize(constants.NORMAL)), pt)
	f.SetSrc(image.NewUniform(color.RGBA{R: 255, G: 238, B: 0, A: 255}))
	pt = freetype.Pt(img.Bounds().Dx()-250, img.Bounds().Dy()-640)
	f.DrawString(fmt.Sprintf("潜伏期人数：%d", pp.GetPeopleSize(constants.SHADOW)), pt)
	f.SetSrc(image.NewUniform(color.RGBA{R: 255, G: 0, B: 0, A: 255}))
	pt = freetype.Pt(img.Bounds().Dx()-250, img.Bounds().Dy()-610)
	f.DrawString(fmt.Sprintf("发病者人数：%d", pp.GetPeopleSize(constants.CONFIRMED)), pt)
	f.SetSrc(image.NewUniform(color.RGBA{R: 72, G: 255, B: 255, A: 255}))
	pt = freetype.Pt(img.Bounds().Dx()-250, img.Bounds().Dy()-580)
	f.DrawString(fmt.Sprintf("已隔离人数：%d", pp.GetPeopleSize(constants.FREEZE)), pt)
	f.SetSrc(image.NewUniform(color.RGBA{R: 0, G: 255, B: 0, A: 255}))
	pt = freetype.Pt(img.Bounds().Dx()-250, img.Bounds().Dy()-550)
	f.DrawString(fmt.Sprintf("空余病床：%d", int(math.Max(float64(constants.BedCount-pp.GetPeopleSize(constants.FREEZE)), 0))), pt)
	f.SetSrc(image.NewUniform(color.RGBA{R: 227, G: 148, B: 118, A: 255}))
	pt = freetype.Pt(img.Bounds().Dx()-250, img.Bounds().Dy()-520)
	f.DrawString(fmt.Sprintf("急需病床：%d", int(math.Max(float64(pp.GetPeopleSize(constants.CONFIRMED)-pp.GetPeopleSize(constants.FREEZE)), 0))), pt)
	f.SetSrc(image.NewUniform(color.RGBA{R: 0, G: 0, B: 0, A: 255}))
	pt = freetype.Pt(img.Bounds().Dx()-250, img.Bounds().Dy()-490)
	f.DrawString(fmt.Sprintf("病死人数：%d", pp.GetPeopleSize(constants.DEATH)), pt)
	f.SetSrc(image.NewUniform(color.RGBA{R: 204, G: 187, B: 204, A: 255}))
	pt = freetype.Pt(img.Bounds().Dx()-250, img.Bounds().Dy()-460)
	f.DrawString(fmt.Sprintf("治愈人数：%d", pp.GetPeopleSize(constants.CURED)), pt)
	f.SetSrc(image.NewUniform(color.White))
	pt = freetype.Pt(img.Bounds().Dx()-250, img.Bounds().Dy()-430)
	f.DrawString(fmt.Sprintf("世界时间（天）：%d", global.WorldTime), pt)

	// 时间+1
	global.WorldTime++

	var b bytes.Buffer
	png.Encode(&b, img)
	return b
}
