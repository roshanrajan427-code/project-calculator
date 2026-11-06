package conversion

import (
	"errors"
	"strconv"
)

func Stringtofloats(strings []string) ([]float64, error) {
 var floats []float64
	for _, Stringval := range strings {
		Floatprice, err := strconv.ParseFloat(Stringval, 64)
		if err != nil {
			return nil , errors.New("error converting ")
		}
		floats=append(floats,Floatprice)
	}
	return floats,nil
}
