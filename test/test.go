package main 

import(
	"fmt"
	//"bohao"
	"os"
	"encoding/csv"
)

 
func main() {
	//Open csv
	f, err := os.Open("out_rab.csv")
	if err != nil{panic(err)}

	r := csv.NewReader(f)
	record, err := r.ReadAll()
	if err != nil{panic(err)}

	fmt.Println(record)

}

/* func main() {
	m1:= bohao.Matrix([]float64{1, 2, 3, 4, 5, 6}, []int{2, 3})
	fmt.Println(m1)
	m2:= bohao.Matrix([]float64{1, 2, 3, 4, 5, 6}, []int{3, 2})
	fmt.Println(m2)

	fmt.Println(bohao.Dot(m1, m2))

	fmt.Println(bohao.Transpose(m1))
} */