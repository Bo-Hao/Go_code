package main 

import(
	"fmt"
	"os"

	"encoding/csv"
)

func main() {
	//create csv file 
	c, err := os.Create("test.csv")
	if err != nil{panic(err)}

	w := csv.NewWriter(c)
	data := [][]string{
		{"1", "中国", "23"},
		{"2", "美国", "23"},
		{"3", "bb", "23"},
		{"4", "bb", "23"},
		{"5", "bb", "23"},
	}
	w.WriteAll(data)
	w.Flush()

	//open csv data
	f, err := os.Open("test.csv")
	if err != nil {panic(err)}

	r  := csv.NewReader(f)
	record, err := r.ReadAll()
	if err != nil {panic(err)}

	fmt.Println(record )

	


	
}