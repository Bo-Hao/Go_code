package main 

import (
    "log"
    //"math/rand"
    "os"
    //"time"

    "github.com/go-echarts/go-echarts/charts"
)

func main() {
	// fake data
	data := make([] []float64, 10)
	for i := 0; i < 10; i++{
		data = append(data, []float64{float64(i), float64(i*2)})
	}


	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(charts.TitleOpts{Title: "Scatter-示例图"})

	scatter.AddYAxis("My points", data)

	f, err := os.Create("scatter.html")
	if err != nil {log.Println(err)}
	scatter.Render(f)

}