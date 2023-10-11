package behavioral

type PaymentStrategy interface {
	Pay(amount int) string
}

type CreditCard struct{}

func (c *CreditCard) Pay(amount int) string {
	return "paid using credit card"
}

type Paypal struct{}

func (p *Paypal) Pay(amount int) string {
	return "paid using paypal"
}

type PaymentProcessor struct {
	strategy PaymentStrategy
}

func (p *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentProcessor) Pay(amount int) string {
	return p.strategy.Pay(amount)
}
