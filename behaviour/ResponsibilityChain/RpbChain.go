package ResponsibilityChain

import "fmt"

type partient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
}

type IDepartment interface {
	handler(p *partient)
	SetNext(i IDepartment)
}

type Department struct {
	next IDepartment
	Name string
}

func (d *Department)handler(p *partient){
		fmt.Println(d.Name+"接收成功")
		if d.next == nil{
			fmt.Println("病人处理完成")
			return
		}
		d.next.handler(p)
}
func (d *Department)SetNext(i IDepartment){
	d.next = i
}

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





