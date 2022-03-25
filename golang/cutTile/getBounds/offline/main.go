package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 栅格切片边框/中心点获取

type GridMeta struct {
	Tiles     string    `json:"tiles,omitempty"`      // "L{z}/R{y}/C{x}"
	ImageType string    `json:"imagetype,omitempty"`  // png
	TileSize  []int64   `json:"tilesize,omitempty"`   // [256, 256]
	Extent    []float64 `json:"extent,omitempty"`     // [13399937.265502144, 4311939.2369373929, 13407191.076071926, 4317988.620193001]
	Center    []float64 `json:"center,omitempty"`     // [13403564.170787035, 4314963.928565197]
	MinZoom   int64     `json:"startlevel,omitempty"` // 最小级别
	MaxZoom   int64     `json:"endlevel,omitempty"`   // 最大级别
}

func reqGET(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("发送请求失败")
		return body, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatal(fmt.Sprintf("发送请求失败, status_code = %v", resp.StatusCode))
		return body, errors.New("发送请求失败")
	}

	body, err = ioutil.ReadAll(resp.Body)
	return body, nil
}

func ReadFromSWD(url string, chLx, chLy, chRx, chRy chan float64, chQuitForSelect chan string) {
	// 读取seaweedfs上的metadata.json文件
	body, err := reqGET(url)
	m := GridMeta{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Fatal(fmt.Sprintf("json格式化失败, url= %v"), url)
		return
	}
	// bounds MinX: 115.756020, MinY: 28.015850, MaxX: 122.834203, MaxY: 34.467777
	// y要对调后:
	// 左: 115.756020, 上: 34.467777, 右: 122.834203, 下: 28.015850
	chLx <- m.Extent[0]
	chLy <- m.Extent[3]
	chRx <- m.Extent[2]
	chRy <- m.Extent[1]
	chQuitForSelect <- url
	return
}


func main() {
	urls := []string{
		"http://192.168.0.219:8888/test_grid/grid666/GF1_PMS2_E116.5_N29.1_20200821_L2E0005004876-MSS2/metadata.json",
		"http://192.168.0.219:8888/test_grid/grid666/GF1B_PMS_E94.8_N43.5_20201102_L1A1227895163-MUX/metadata.json",
	}
	chLx := make(chan float64, 0)
	chLy := make(chan float64, 0)
	chRx := make(chan float64, 0)
	chRy := make(chan float64, 0)
	chQuitForSelect := make(chan string, len(urls))
	quitTag := make([]string, 0)
	chQuit := make(chan bool)
	for _, url := range(urls) {
		go ReadFromSWD(url, chLx, chLy, chRx, chRy, chQuitForSelect)
	}
	var minX, minY, maxX, maxY float64 // 外边框
	go func() {
		for {
			select {
			case lx := <- chLx:
				if lx < minX {
					minX = lx
				} else if minX == 0 {
					minX = lx
				}
			case ly := <- chLy:
				if ly < minY {
					minY = ly
				} else if minY == 0 {
					minY = ly
				}
			case rx := <- chRx:
				if maxX < rx {
					maxX = rx
				}
			case ry := <- chRy:
				if maxY < ry {
					maxY = ry
				}
			case q := <- chQuitForSelect:
				fmt.Printf("完成json读取, url = %v \n", q)
				quitTag = append(quitTag, q)
				if len(quitTag) == len(urls) {
					chQuit <- true
				}
			}
		}
	}()
	<- chQuit
	fmt.Println("结束啦")
	fmt.Printf("minX= %v, minY= %v, maxX= %v, maxY= %v \n", minX, minY, maxX, maxY)
	bounds := []float64{minX, minY, maxX, maxY}
	center := []float64{(minX + maxX) / 2, (minY + maxY) / 2}
	fmt.Printf("bounds = %v \n", bounds)
	fmt.Printf("center = %v \n", center)
}
