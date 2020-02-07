package main 

import(
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
	bar := charts.NewBar()
    bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
    bar.AddXAxis(nameItems).
        AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
		
	line := charts.NewLine()
	line.AddXAxis(nameItems).AddYAxis("商家A", randInt())

	scatter := charts.NewScatter()
    scatter.AddXAxis(nameItems).
        AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())

	es := charts.NewEffectScatter()
	es.AddXAxis(nameItems).AddYAxis("es1", randInt())

	p := charts.NewPage()
	p.Add(
		bar, 
		line,
		scatter, 
		es, 
	)

	f, err := os.Create("page.html")
	if err != nil{log.Fatal("error")}

	p.Render(f)

		
}
	