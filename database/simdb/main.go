package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/Bo-Hao/bohao"
	"github.com/sonyarouje/simdb/db"
)


type ResultRecord struct {
	Id     int       `json:"id"`
	TrainX []float64 `json:"trainx"`
	TrainY []float64 `json:"trainy"`
	TestX  []float64 `json:"testx"`
	TestY  []float64 `json:"testy"`
	MB     float64   `json:"mb"`
	MR     float64   `json:"mr"`
	MAB    float64   `json:"mab"`
	RMSE   float64   `json:"rmse"`
	SDR    float64   `json:"sdr"`
}

func (c ResultRecord) ID() (jsonField string, value interface{}) {
	value = c.Id
	jsonField = "custid"
	return
}

func test2() {
	f, _ := os.Open("/Users/pengbohao/Go_code/database/simdb/Inv_1_.csv")
	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	record, err := r.ReadAll()
	if err != nil {
		panic(nil)
	}

	fmt.Println("starting....")
	driver, err := db.New("dbs")
	if err != nil {
		panic(err)
	}

	n := 0
	for i := 0; i < len(record); i++ {
		if record[i][0] == "MB" {
			R := ResultRecord{}

			data := bohao.ConvSliceFromStr2Float(record[i-4 : i])
			R.Id = n
			R.TrainX = data[0]
			R.TrainY = data[1]
			R.TestX = data[2]
			R.TestY = data[3]
			R.MB, _ = strconv.ParseFloat(record[i][1], 64)
			R.MR, _ = strconv.ParseFloat(record[i][3], 64)
			R.MAB, _ = strconv.ParseFloat(record[i][5], 64)
			R.RMSE, _ = strconv.ParseFloat(record[i][7], 64)
			R.SDR, _ = strconv.ParseFloat(record[i][9], 64)
			n += 1

			err = driver.Insert(R)
			if err != nil {
				panic(err)
			}
		}
	}
	var results []ResultRecord
	driver.Open(ResultRecord{}).Get().AsEntity(&results)
	fmt.Println(results[:3])
}
func main() {
	driver, err := db.New("dbs")
	if err != nil {
		panic(err)
	}

	var results []ResultRecord
	driver.Open(ResultRecord{}).Get().AsEntity(&results)
	fmt.Println(len(results))

}
