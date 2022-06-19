package bridge

import "fmt"

type Computer interface {
	print()
	setPrinter(Printer)
}

func GetComputer(name string,p Printer)(Computer,error){
	switch name {
	case "mac":
		return &mac{printer:p},nil
	case "windows":
		return &windows{printer:p},nil
	default:
		return nil,fmt.Errorf("错误类型")
	}
}

type mac struct {
	printer Printer
}

func (m *mac) print() {
	fmt.Println("Print request for mac")
	m.printer.printFile()
}

func (m *mac) setPrinter(p Printer) {
	m.printer = p
}

type windows struct {
	printer Printer
}

func (w *windows) print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(p Printer) {
	w.printer = p
}

func GetPrinter(name string)(Printer,error){
	switch name {
	case "epson":
		return &epson{},nil
	case "hp":
		return &hp{},nil
	default:
		return nil,fmt.Errorf("错误的类型")
	}
}

type Printer interface {
	printFile()
}

type epson struct {
}

func (p *epson) printFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type hp struct {
}

func (p *hp) printFile() {
	fmt.Println("Printing by a HP Printer")
}