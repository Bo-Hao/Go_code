package main 

import(
	"os"
	"log"
	"github.com/go-echarts/go-echarts/charts"
	
)


func main() {
	kline := charts.NewKLine()

	x := make([]string, 0)
	y := make([][4]float32, 0)
	for i := 0; i < len(kd); i++ {
		x = append(x, kd[i].date)
		y = append(y, kd[i].data)
	}
	
	kline.AddXAxis(x).AddYAxis("kline", y)
	kline.SetGlobalOptions(
		charts.TitleOpts{Title: "Kline-示例图"},
		charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		charts.DataZoomOpts{Type:"inside", XAxisIndex: []int{0}, Start: 50, End: 100},
    	charts.DataZoomOpts{Type:"slider", XAxisIndex: []int{0}, Start: 50, End: 100},
	)

	f, err := os.Create("kline.html")
	if err != nil{log.Fatal("error")}

	kline.Render(f)
}