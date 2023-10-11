package structural

import "fmt"

type ClimatService struct{}

func (c *ClimatService) GetTemperatureAndHumidity() string {
	return "25.5, 50.5"
}

type NewService struct{}

func (n *NewService) GetNews() string {
	return "Headline: This is an important news!"
}

type StockService struct{}

func (s *StockService) GetStock() string {
	return "Stock: 1000"
}

type Facade struct {
	climatService *ClimatService
	newService    *NewService
	stockService  *StockService
}

func NewFacade() *Facade {
	return &Facade{
		climatService: &ClimatService{},
		newService:    &NewService{},
		stockService:  &StockService{},
	}
}

func (f *Facade) GetAggregatedData() string {
	w := f.climatService.GetTemperatureAndHumidity()
	n := f.newService.GetNews()
	s := f.stockService.GetStock()
	return fmt.Sprintf("%s\n%s\n%s", w, n, s)

}
