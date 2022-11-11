package iterator

type people struct {
	Name string
	Age  int
}

type iterator interface {
	hasNext() bool
	GetNext() *people
}

type PeopleIterator struct {
	index  int
	people []*people
}

func (p *PeopleIterator) hasNext() bool {
	if p.index < len(p.people) {
		return true
	}
	return false
}

func (p *PeopleIterator) GetNext() *people {
	if p.hasNext() {
		p.index = p.index + 1
		return p.people[p.index-1]
	}
	return nil
}

func (p *PeopleIterator) SetIterator(p1 []*people) {
	p.people = p1
}
