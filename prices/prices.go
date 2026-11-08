package prices

import (
	"calci/conversion"
	"calci/filemanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

// for reading from a file
func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.FileRead()
	if err != nil {
		fmt.Println(err)
		return
	}

	Prices, err := conversion.Stringtofloats(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = Prices

}

// method function
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedprice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedprice)
	}
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
}

// constructor function
func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
