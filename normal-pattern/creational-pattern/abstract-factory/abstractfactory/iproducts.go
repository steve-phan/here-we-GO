package abstractfactory

import "fmt"

// import "fmt"

type propsBasic struct {
	logo  string
	price int
}

type methodsBasic interface {
	// setLogo(logo string)
	// setPrice(price int)
	GetLogo() string
	GetPrice() int
}

// Setter and Getter for both Phone and Table

func (p *propsBasic) setLogo(logo string) {
	fmt.Println("Creating ....")
	p.logo = logo
}

func (p *propsBasic) setPrice(price int) {
	fmt.Println("Creating ....")
	p.price = price
}

func (p *propsBasic) GetLogo() string {
	return p.logo
}

func (p *propsBasic) GetPrice() int {
	return p.price
}
