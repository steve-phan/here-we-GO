package abstractfactory

type IPhone interface {
	methodsBasic
	GetMarket() string
}

type ITable interface {
	methodsBasic
	IsSupport5G() bool
}

type ApplePhone struct {
	propsBasic
	market string
}

type AppleTable struct {
	propsBasic
	is5G bool
}

type Apple struct {
}

func NewApple() *Apple {
	return &Apple{}
}

func (a *Apple) MakePhone() IPhone {
	return &ApplePhone{
		propsBasic: propsBasic{
			logo:  "APPLE",
			price: 900,
		},
		market: "DE",
	}
}

func (a *Apple) MakeTable() ITable {
	return &AppleTable{
		propsBasic: propsBasic{
			logo:  "APPLE",
			price: 750,
		},
		is5G: true,
	}
}

// A method get customize property of Phone
func (p *ApplePhone) GetMarket() string {
	return p.market
}

// // A method get customize property of Table
func (t *AppleTable) IsSupport5G() bool {
	return t.is5G
}
