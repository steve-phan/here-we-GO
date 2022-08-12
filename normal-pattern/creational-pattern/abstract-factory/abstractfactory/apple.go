package abstractfactory

type ApplePhone struct {
	Phone
}

type Apple struct {
}

// type AppleTable struct {
// 	Table
// }

func NewApple() *Apple {
	return &Apple{}
}

func (a *Apple) MakePhone() IPhone {
	return &ApplePhone{
		Phone: Phone{
			logo:  "APPLE",
			price: 900,
		},
	}
}
