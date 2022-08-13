package chanres

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) Excute(p *Patient) {
	// Check if the patient is registered or not
	if p.isRegistered {
		fmt.Println("Patient is already registered")
		r.next.Excute(p)
		return
	}
	fmt.Println(("Patient is registering..."))
	p.isRegistered = true
	r.next.Excute(p)
}

func (r *Reception) SetNext(next Department) {
	// Where is the next department for the patient
	r.next = next
}
