package ResponsibilityChain

import "testing"

func TestRpbChain(t *testing.T) {

	//Set next for medical department
	medical := &Medical{Department{Name: "药房"}}
	//Set next for doctor department
	doctor := &Doctor{Department{Name: "医生"}}
	//Set next for reception department
	reception := &Receiption{Department{Name: "前台"}}
	patient := &partient{name: "张三"}
	//Patient visiting
	rpb := CreRpbManager().AddIter("通用", CreateIter(reception, doctor, medical))
	if v := rpb.SelectIter("通用"); v != nil {
		v.GOGOGO(patient)
		return
	}
	t.Fatal("没有相对应的病症治疗")
}
