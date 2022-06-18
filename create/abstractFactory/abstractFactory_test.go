package abstractFactory

import "testing"

func TestGetSportFactory(t *testing.T) {
	adidasFactory,_ := GetSportFactory("adidas")
	printShoeDetails(adidasFactory.makeShoe())
	printShirtDetails(adidasFactory.makeShirt())
}