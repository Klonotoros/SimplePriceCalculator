package main

import (
	"fmt"
	"simple-price-calculator/cmdmanager"
	"simple-price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {

		//fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPrices(cmdm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("couldn't process data")
			fmt.Println(err)
		}

	}

}
