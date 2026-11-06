package main

import (
        "calci/prices")

func main() {
TaxRates:=[]float64{0,0.7,0.1,0.15}
for _,taxRate:=range TaxRates{
	PriceJob:=prices.NewTaxIncludedPriceJob(taxRate)
	PriceJob.Process()
}
}




