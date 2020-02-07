package main

import (
    "log"
    "math/rand"
    "os"
    

    "github.com/go-echarts/go-echarts/charts"
)

var rangeColor = []string{
    "#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
    "#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}

func genScatter3dData() [][3]int {
    data := make([][3]int, 0)
    for i := 0; i < 80; i++ {
        data = append(data, [3]int{
            int(rand.Int63()) % 100,
            int(rand.Int63()) % 100,
            int(rand.Int63()) % 100,
        })
    }
    return data
}

func main() {
    scatter3d := charts.NewScatter3D()
    
    //set global option 
	scatter3d.SetGlobalOptions(
		charts.TitleOpts{Title: "Scatter3D-示例图"},
		charts.VisualMapOpts{
			Calculable: true, //根據value的選取以及顏色漸層
		    InRange:    charts.VMInRange{Color: rangeColor},
            Max:        100,

		},
		charts.Grid3DOpts{BoxDepth: 80, BoxWidth: 200},
    )
    

	scatter3d.AddZAxis("scatter3d", genScatter3dData())
	
	f, err := os.Create("scatter3D.html")
    if err != nil {
        log.Println(err)}
	scatter3d.Render(f)
	
}