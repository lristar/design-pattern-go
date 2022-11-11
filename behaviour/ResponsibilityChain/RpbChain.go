package ResponsibilityChain

import "fmt"

// 换另一种模式,不应该是死绑定的,需要一个责任链管理者,例如医院
type RpbManager struct {
	Iters map[string]*RpbIter
}

func CreRpbManager() *RpbManager {
	return &RpbManager{Iters: make(map[string]*RpbIter, 0)}
}

func (r *RpbManager) AddIter(key string, iter *RpbIter) *RpbManager {
	if _, ok := r.Iters[key]; !ok {
		r.Iters[key] = iter
	}
	return r
}

func (r *RpbManager) SelectIter(key string) *RpbIter {
	if v, ok := r.Iters[key]; ok {
		return v
	}
	return nil
}

type RpbIter struct {
	IDepartments DepartmentIterator
}

func CreateIter(d ...IDepartment) *RpbIter {
	return &RpbIter{
		IDepartments: DepartmentIterator{index: -1, departs: d},
	}
}

func (r *RpbIter) GOGOGO(p *partient) {
	for r.IDepartments.hasNext() {
		r.IDepartments.GetNext().handler(p)
	}
}

type partient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
}

type IDepartment interface {
	handler(p *partient)
}

type Department struct {
	next IDepartment
	Name string
}

func (d *Department) handler(p *partient) {
	fmt.Println(d.Name + "接收成功")
	if d.next == nil {
		fmt.Printf("病人%s处理完成\n", p.name)
		return
	}
}

//type IDepartment interface {
//	handler(p *partient)
//	SetNext(i IDepartment)
//}
//
//
//func (d *Department) handler(p *partient) {
//	fmt.Println(d.Name + "接收成功")
//	if d.next == nil {
//		fmt.Println("病人处理完成")
//		return
//	}
//	d.next.handler(p)
//}
//func (d *Department) SetNext(i IDepartment) {
//	d.next = i
//}

// 前台接收者
type Receiption struct {
	Department
}

// 医生
type Doctor struct {
	Department
}

// 药房
type Medical struct {
	Department
}

// Iterator 做一个迭代器
type Iterator interface {
	hasNext() bool
	GetNext() IDepartment
}

type DepartmentIterator struct {
	index         int
	departs       []IDepartment
	globalConfigs map[string]string
}

func (d *DepartmentIterator) hasNext() (is bool) {
	if d.index >= len(d.departs)-1 {
		return
	}
	return true
}

func (d *DepartmentIterator) GetNext() IDepartment {
	if d.hasNext() {
		d.index++
		return d.departs[d.index]
	}
	return nil
}

func (d *DepartmentIterator) SetIter(depart []IDepartment) {
	d.departs = depart
}
