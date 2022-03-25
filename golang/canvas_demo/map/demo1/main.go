package main

import (
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
	"image/color"
)

func main() {
	c := canvas.New(256, 256)
	ctx := canvas.NewContext(c)
	ctx.SetFillColor(color.RGBA{255, 0, 0, 255})   // 填充颜色
	ctx.SetStrokeColor(color.RGBA{0, 0, 255, 255})  // 边框颜色
	ctx.SetStrokeWidth(50)  // 边框宽度
	//ctx.ResetView()
	ctx.DrawPath(0.0, 0.0, canvas.Rectangle(ctx.Width(), ctx.Height()))
	renderers.Write("/home/ean/Desktop/demo.png", c)
}
