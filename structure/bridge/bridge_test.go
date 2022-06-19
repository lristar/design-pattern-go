package bridge

import "testing"

func TestBridge(t *testing.T) {
	printer,err:=GetPrinter("epson")
	if err!=nil{
		t.Fatal(err)
	}
	cpm,err:=GetComputer("mac",printer)
	if err!=nil{
		t.Fatal(err)
	}
	cpm.print()
}
