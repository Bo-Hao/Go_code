package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/charts"
)

var nameItems = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}
var seed = rand.NewSource(time.Now().UnixNano())

func randInt() []int {
	cnt := len(nameItems)
	r := make([]int, 0)
	for i := 0; i < cnt; i++ {
		r = append(r, int(seed.Int63())%50)
	}
	return r
}

func main() {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Line-显示 Label"})
	line.AddXAxis(nameItems)
	line.AddYAxis("商家A", randInt(),
		charts.LabelTextOpts{Show: true},
		charts.LineOpts{Smooth: true},
	)

	f, err := os.Create("line.html")
	if err != nil {
		log.Fatal("error")
	}

	line.Render(f)

}
