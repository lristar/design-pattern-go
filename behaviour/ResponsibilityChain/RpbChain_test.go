package ResponsibilityChain

import "testing"

func TestRpbChain(t *testing.T) {

	//Set next for medical department
	medical := &Medical{Department{Name:"药房"}}
	//Set next for doctor department
	doctor := &Doctor{Department{Name:"医生"}}
	doctor.SetNext(medical)
	//Set next for reception department
	reception := &Receiption{Department{Name:"前台"}}
	reception.SetNext(doctor)
	patient := &partient{name: "张三"}
	//Patient visiting
	reception.handler(patient)
}