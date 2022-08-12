package abstractfactory

type PropsBasic struct {
	logo  string
	price int
}

type MethodsBasic interface {
	setLogo(logo string)
	setPrice(price int)
	GetLogo() string
	GetPrice() int
}

type Phone struct {
	logo  string
	price int
}
type Table struct {
	PropsBasic
}

type IPhone interface {
	MethodsBasic
}

type ITable interface {
	MethodsBasic
}

// Settter and Getter for phone

func (p *Phone) setLogo(logo string) {
	p.logo = logo
}

func (p *Phone) setPrice(price int) {
	p.price = price
}

func (p *Phone) GetLogo() string {
	return p.logo
}

func (p *Phone) GetPrice() int {
	return p.price
}

// Settter and Getter for table

// func (p *Table) setLogo(logo string) {
// 	p.logo = logo
// }

// func (p *Table) setPrice(price int) {
// 	p.price = price
// }

// func (p *Table) GetLogo() string {
// 	return p.logo
// }

// func (p *Table) GetPrice() int {
// 	return p.price
// }
