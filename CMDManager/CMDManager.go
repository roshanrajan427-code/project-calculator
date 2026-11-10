package CMDManager

import "fmt"

type CMDManager struct {
}

func (cmd CMDManager) FileRead() ([]string, error) {
fmt.Println("Please enter ENTER  after every price")
var prices []string
for {
	var price string
	fmt.Print("price:")
	fmt.Scan(&price)
	if price=="0"{
		break;
	}
    prices = append(prices, price)
}
return  prices,nil
}
func (cmd CMDManager) WriteResult(data any) error {
	fmt.Print(data)
	return nil
}