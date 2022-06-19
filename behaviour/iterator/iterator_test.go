package iterator

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	p1 :=&people{Name:"张三",Age:18}
	p2 :=&people{Name:"李四",Age:19}
	p3 :=&people{Name:"王五",Age:18}
	p :=&PeopleIterator{people:[]*people{p1,p2,p3}}
	for p.hasNext(){
		pp :=p.GetNext()
		fmt.Printf("User is %+v\n", *pp)
	}
}