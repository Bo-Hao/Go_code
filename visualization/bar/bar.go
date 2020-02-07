package main

import(
	"log"
	"os"

	"github.com/chenjiandongx/go-echarts/charts"
)

func main() {
	nameItems := []string{"1", "2", "3", "4", "5", "6"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar"})
	bar.AddXAxis(nameItems)
	bar.AddYAxis("A", []int{20, 30, 40, 10, 24, 36,})
	bar.AddYAxis("B", []int{35, 14, 25, 60, 44, 23,})
	f, err := os.Create("bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(f)
}