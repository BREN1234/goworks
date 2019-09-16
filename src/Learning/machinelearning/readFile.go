package machinelearning

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
)

func fun() {
	df, err := base.ParseCSVToInstances("D:/Datasets/AnalyticVidyaDatasets/stockdata/stock.csv", false)
	if err != nil {
		panic(err)
	}
	fmt.Println(df)
}
