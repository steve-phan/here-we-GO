package chanres

import "fmt"

type Medical struct {
	next Department
}

func (r *Medical) Excute(p *Patient) {
	// Check if the medicice is taken by the patient or not
	if p.isMedicineTaken {
		fmt.Println("Patient is already checked by the doctor")
		r.next.Excute(p)
		return
	}
	fmt.Println(("Patient is checking by the doctor..."))
	p.isMedicineTaken = true
	r.next.Excute(p)
}

func (r *Medical) SetNext(next Department) {
	// Where is the next department for the patient
	r.next = next
}
