package prices

import (
	"bufio"
	"calci/conversion"
	"fmt"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

// for reading from a file
func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("error reading file")
		fmt.Println(err)
		file.Close()
		return
	}
	Prices,err := conversion.Stringtofloats(lines)
		if err != nil {
			fmt.Println(err)
			file.Close()
			return
		}
		
	
	job.InputPrices = Prices
  file.Close()
  }

// method function
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedprice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedprice)
	}
	fmt.Println(result)
}

// constructor function
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
