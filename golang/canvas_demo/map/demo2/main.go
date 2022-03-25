package main

import (
	"canvas_demo/map/demo2/model"
	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/encoding/wkb"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
	"log"
)

func demo() {
	// 查询数据
	a := model.Adcode{}
	geobytes := a.SearchFirst()
	g, err := wkb.DecodeBytes(geobytes)
	if err != nil {
		log.Fatal("wkb.DecodeBytes error, err: ", err)
	}
	// 转换结构体
	gm := g.(geom.MultiPolygon)
	mpg := gm.Polygons()
	switch g.(type) {
	case geom.MultiPolygon:

		for _, poly := range mpg {
			for _, ring := range poly {
				if len(ring) == 0 {
					continue
				}
				p := &canvas.Path{}
				p.MoveTo(ring[0][0], ring[0][1])
				for _, point := range ring {
					p.LineTo(point[0], point[1])
				}
				p.Close()
			}
		}
	}

	// 创建画布
	c := canvas.New(256, 256)
	ctx := canvas.NewContext(c)
	draw(ctx, geobytes)
	renderers.Write("/home/ean/Desktop/out_123213.png", c, canvas.DPMM(8.0))
}

func draw(c *canvas.Context, geobytes []byte) {
	// 转换结构体
	// 本例明确知道数据为multipolygon
	g, err := wkb.DecodeBytes(geobytes)
	if err != nil {
		log.Fatal("wkb decode error, err: ", err)
	}
	mpg := g.(geom.MultiPolygon).Polygons()
	var xmin, xmax, ymin, ymax float64
	for _, poly := range mpg {
		for _, ring := range poly {
			if len(ring) == 0 {
				continue
			}
			p := &canvas.Path{}
			p.MoveTo(ring[0][0], ring[0][1])
			for _, point := range ring {
				if point[0] > xmax {
					xmax = point[0]
				}
				if point[0] < xmin {
					xmin = point[0]
				}
				if point[1] > ymax {
					ymax = point[1]
				}
				if point[1] < ymin {
					ymin = point[1]
				}
				p.LineTo(point[0], point[1])
			}
			p.Close()
		}
	}
	xscale := 100.0 / (xmax - xmin)
	yscale := 100.0 / (ymax - ymin)
	c.SetView(canvas.Identity.Translate(0.0, 0.0).Scale(xscale, yscale).Translate(-xmin, -ymin))
	c.SetStrokeWidth(0.5)
	c.ResetView()
	c.DrawPath(0.0, 0.0, canvas.Rectangle(c.Width(), c.Height()))
}

func main() {
	err := model.ConnectToDB(model.DSN)
	if err != nil {
		log.Fatal("connect db error, err: ", err)
	}
	demo()
}
