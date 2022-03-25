package model

import (
	"fmt"
	"log"
)

type Adcode struct {

}

func (a *Adcode) SearchFirst() []byte {
	sql := fmt.Sprintf(`SELECT ST_AsBinary(st_transform(the_geom, 4326)) AS the_geom FROM adcode WHERE fid = 1`)
	rows, err := DB.Raw(sql).Rows()
	if err != nil {
		log.Fatal("search sql error, err: ", err)
	}
	geobytes := make([]byte, 0)
	for rows.Next() {
		rows.Scan(&geobytes)
	}
	return geobytes
}
