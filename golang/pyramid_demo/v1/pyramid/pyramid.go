package pyramid

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/encoding/mvt"
	"github.com/go-spatial/geom/encoding/wkb"
	"github.com/go-spatial/geom/slippy"
	"log"
	"pyramid_demo/model"
	"pyramid_demo/tegola/convert"
	"pyramid_demo/tegola/maths/validate"
)

func Pyramid() {
	// 1.获取数据外包矩形
	var strb string
	model.DB.Raw("SELECT st_asbinary(st_extent(the_geom)) FROM adcode_test").Scan(&strb)
	b := []byte(strb)
	gBound, err := wkb.DecodeBytes(b)
	if err != nil {
		log.Fatalln("byte can not decode")
	}
	ex, err := geom.NewExtentFromGeometry(gBound)
	if err != nil {
		log.Fatalln("get extent from geometry error: ", err)
	}
	// 2.创建相应的金字塔索引表
	sqls := []string{
		`create table if not exists adcode_test_pyramid (id bigserial primary key,x integer,y integer,z integer,pyramid jsonb);`,
		`create unique index if not exists idx_x_y_z on adcode_test_pyramid (x, y, z)`,
	}
	for _, s := range sqls {
		model.DB.Debug().Exec(s)
	}
	// 3.获取某级所有瓦片编号
	tiles := make([]slippy.Tile, 0)
	for i := 0; i <=13 ; i++ {
		t := slippy.FromBounds(ex, uint(i))
		tiles = append(tiles, t...)
	}
	// 4.遍历瓦片数组, 与数据库交互获取geometry字段
	for _, tile := range tiles {
		tile.X = 25
		tile.Y = 26
		tile.Z = 5
		extent := tile.Extent3857().ExpandBy(slippy.Pixels2Webs(tile.Z, 64))
		fmt.Println(extent)
		sql := fmt.Sprintf(`SELECT fid, ST_AsBinary(st_transform(the_geom, 3857)) AS the_geom FROM adcode_test
WHERE the_geom && st_transform(ST_MakeEnvelope(%v,%v,%v,%v, 3857), 4326)`, extent[0], extent[1], extent[2], extent[3])
		rows, err := model.DB.Debug().Raw(sql).Rows()
		if err != nil {
			log.Fatalln("search db error, err: ", err)
		}
		var (
			geobytes []byte
			fid uint
		)
		pyramidMap := make(map[uint]interface{}, 0)
		clipRegion := geom.NewExtent([2]float64{-64, -64}, [2]float64{4160, 4160})
		for rows.Next() {
			rows.Scan(&fid, &geobytes)
			geometry, err := wkb.DecodeBytes(geobytes)
			if err != nil {
				log.Fatalln("wkb DecodeBytes error, err: ", err)
			}
			geo := mvt.PrepareGeo(geometry, tile.Extent3857(), float64(mvt.DefaultExtent))

			ctx := context.Background()
			sg, err := convert.ToTegola(geo)
			tegolaGeo, err := validate.CleanGeometry(ctx, sg, clipRegion)
			geo, err = convert.ToGeom(tegolaGeo)
			b, err := wkb.EncodeBytes(geo)
			if err != nil {
				log.Fatalln("can nao encode geo to byte, err: ", err)
			}
			pyramidMap[fid] = b
		}
		pb, err := json.Marshal(pyramidMap)
		if err != nil {
			log.Fatalln("json marshal error, err: ", err)
		}
		// 5.组建金字塔并往金字塔索引表中插入数据
		model.DB.Exec(`insert into adcode_test_pyramid(x, y, z, pyramid) values (?, ?, ?, ?)`, tile.X, tile.Y, tile.Z, pb)
	}

}
