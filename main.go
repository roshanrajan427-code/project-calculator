package main

import (
	"calci/filemanager"
	"calci/prices"
	"fmt"
)

func main() {
	TaxRates := []float64{0, 0.7, 0.1, 0.15}
	for _, taxRate := range TaxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		PriceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		PriceJob.Process()
	}
}
