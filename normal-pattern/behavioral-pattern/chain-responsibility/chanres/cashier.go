package chanres

import "fmt"

type Cashier struct {
	next Department
}

func (r *Cashier) Excute(p *Patient) {
	// Check if the patient is checked by the doctor or not
	if p.isPaid {
		fmt.Println("Patient is already paid")
		r.next.Excute(p)
		return
	}
	fmt.Println(("Patient is paying..."))
	p.isPaid = true

}

func (r *Cashier) SetNext(next Department) {
	// Where is the next department for the patient
	r.next = next // do somethings else ???
	// fmt.Println("DONE")
}
