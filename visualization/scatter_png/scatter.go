package main

import (
	"gonum.org/v1/plot/vg/draw"
	//"fmt"
	"math/rand"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func line_plot_1() {
	pts := plotter.XYs{}

	for i := 0; i < 10; i++{
		pts = append(pts, plotter.XY{
			X: float64(i), 
			Y: 0.7 * float64(i) + 1, 
		})
	}

	plt, err := plot.New()
	if err != nil{panic(err)}

	plt.Y.Min, plt.X.Min, plt.Y.Max, plt.X.Max = 0, 0, 10, 10

	err = plotutil.AddLines(plt, "line1", pts,)
	if err != nil {panic(err)}

	err = plt.Save(5*vg.Inch, 5*vg.Inch, "01-line.png") 
	if err != nil {panic(err)}
}

func line_plot_2() {
	var a, b float64 = 0.7, 3
	points1 := plotter.XYs{}
	points2 := plotter.XYs{}

	for i := 0; i <= 10; i++ {
		points1 = append(points1, plotter.XY{
			X: float64(i),
			Y: a*float64(i) + b,
		})
		points2 = append(points2, plotter.XY{
			X: float64(i),
			Y: a*float64(i) + b + (2*rand.Float64() - 1),
		})
	}

	plt, err := plot.New()
	if err != nil{panic(err)}

	plt.Y.Min, plt.X.Min, plt.Y.Max, plt.X.Max = 0, 0, 10, 10

	err = plotutil.AddLinePoints(plt, "line1", points1, "line2", points2,)
	if err != nil {panic(err)}

	err = plt.Save(5*vg.Inch, 5*vg.Inch, "02-line.png") 
	if err != nil {panic(err)}
}
func LeastSquares(points plotter.XYs) (a, b float64) {
	var xSum, ySum float64
	for _, point := range points {
		xSum += point.X
		ySum += point.Y
	}
	xAvg, yAvg := xSum/float64(points.Len()), ySum/float64(points.Len())

	var xySum, xxSum float64
	for _, point := range points {
		xySum += (point.X - xAvg) * (point.Y - yAvg)
		xxSum += (point.X - xAvg) * (point.X - xAvg)
	}

	a = xySum / xxSum
	b = yAvg - a*xAvg
	return
}

func regression_line() {
	var a, b float64 = 0.7, 3
	points1 := plotter.XYs{}
	points2 := plotter.XYs{}

	for i := 0; i <= 10; i++ {
		points1 = append(points1, plotter.XY{
			X: float64(i),
			Y: a*float64(i) + b,
		})
		points2 = append(points2, plotter.XY{
			X: float64(i),
			Y: a*float64(i) + b + (2*rand.Float64() - 1),
		})
	}


	fa, fb := LeastSquares(points2)
	points3 := plotter.XYs{}
	for i := 0; i <= 10; i++ {
		points3 = append(points3, plotter.XY{
			X: float64(i),
			Y: fa*float64(i) + fb,
		})
	}

	plt, err := plot.New()
	if err != nil{panic(err)}

	plt.Y.Min, plt.X.Min, plt.Y.Max, plt.X.Max = 0, 0, 10, 10

	if err := plotutil.AddLinePoints(plt,
		"line1", points1,
		"line2", points2,
		"line3", points3,
	); err != nil {
		panic(err)
	}

	err = plt.Save(5*vg.Inch, 5*vg.Inch, "03-line.png") 
	if err != nil {panic(err)}
	
}

func scatter_plot() {
	points := plotter.XYs{}
	for i := 0; i < 10; i++ {
		points = append(points, plotter.XY{
			X: 100 * rand.Float64(),
			Y: 100 * rand.Float64(),
		})
	}

	scatter, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}
	scatter.Shape = draw.CircleGlyph{}

	plt, err := plot.New()
	if err != nil {
		panic(err)
	}
	plt.Y.Min, plt.X.Min, plt.Y.Max, plt.X.Max = 0, 0, 100, 100

	plt.Add(scatter)

	err = plt.Save(5*vg.Inch, 5*vg.Inch, "04-scatter.png") 
	if err != nil {panic(err)}


}

func main() {
	line_plot_1()	//01 line
	line_plot_2()	//02 line
	regression_line() //03 line 
	scatter_plot()

}
