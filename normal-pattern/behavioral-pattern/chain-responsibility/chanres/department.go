package chanres

type Department interface {
	// need pointer to Patient to track status of this patient
	Excute(p *Patient)
	SetNext(d Department)
}
