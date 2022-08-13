package chanres

import "fmt"

type Doctor struct {
	next Department
}

func (r *Doctor) Excute(p *Patient) {
	// Check if the patient is checked by the doctor or not
	if p.isDoctorChecked {
		fmt.Println("Patient is already checked by the doctor")
		r.next.Excute(p)
		return
	}
	fmt.Println(("Patient is checking by the doctor..."))
	p.isDoctorChecked = true
	r.next.Excute(p)
}

func (r *Doctor) SetNext(next Department) {
	// Where is the next department for the patient
	r.next = next
}
