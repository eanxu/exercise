package pyramid

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/encoding/mvt"
	"github.com/go-spatial/geom/encoding/wkb"
	"github.com/go-spatial/geom/slippy"
	"github.com/lib/pq"
	"github.com/panjf2000/ants/v2"
	_ "github.com/panjf2000/ants/v2"
	"log"
	"pyramid_demo/model"
	"pyramid_demo/tegola/convert"
	"pyramid_demo/tegola/maths/validate"
	"pyramid_demo/utils"
	"sync"
)

func Pyramid() error {
	// 1.获取数据外包矩形
	var strb string
	model.DB.Raw("SELECT st_asbinary(st_extent(the_geom)) FROM adcode_test").Scan(&strb)
	b := []byte(strb)
	ex, err := utils.GetBoundGeometry(b)
	if err != nil {
		return err
	}

	// 2.创建相应的金字塔索引表
	sqls := []string{
		`create table if not exists adcode_test_pyramid_v2 (x integer,y integer,z integer, ids integer[], pyramid text);`,
		`create index if not exists idx_x_y_z on adcode_test_pyramid (x, y, z)`,
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

	// 4. 切片入库
	var wg sync.WaitGroup

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(50, func(i interface{}) {
		job(i.(slippy.Tile))
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for _, tile := range tiles {
		wg.Add(1)
		_ = p.Invoke(tile)
	}
	wg.Wait()

	return nil
}

func job(tile slippy.Tile) {
	extent := tile.Extent3857().ExpandBy(slippy.Pixels2Webs(tile.Z, 64))
	sql := fmt.Sprintf(`SELECT fid, ST_AsBinary(st_transform(the_geom, 3857)) AS the_geom FROM adcode_test
WHERE the_geom && st_transform(ST_MakeEnvelope(%v,%v,%v,%v, 3857), 4326)`, extent[0], extent[1], extent[2], extent[3])
	rows, err := model.DB.Raw(sql).Rows()
	if err != nil {
		log.Fatalln("search db error, err: ", err)
	}
	var (
		geobytes []byte
		fid int64
	)
	pyramidMap := make(map[int64]interface{}, 0)
	fids := make(pq.Int64Array, 0)
	clipRegion := geom.NewExtent([2]float64{-64, -64}, [2]float64{4160, 4160})
	for rows.Next() {
		rows.Scan(&fid, &geobytes)
		fids = append(fids, fid)
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

	if len(fids) == 0 {
		// 判空操作
		return
	}

	pb, err := json.Marshal(pyramidMap)
	if err != nil {
		log.Fatalln("json marshal error, err: ", err)
	}
	// 5.组建金字塔并往金字塔索引表中插入数据
	// pbStr := string(pb)
	model.DB.Exec(`insert into adcode_test_pyramid_v2 (x, y, z, ids, pyramid) values (?, ?, ?, ?, ?)`, tile.X, tile.Y, tile.Z, fids, pb)
}
