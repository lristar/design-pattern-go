package abstractFactory

import "fmt"

type ISportFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}
type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

func GetSportFactory(ty string)(ISportFactory,error){
	switch ty {
	case "adidas":
		return &Adidas{},nil
	case "nike":
		return &Nike{},nil
	default:
		return nil,fmt.Errorf("类型错了")
	}
}

type Adidas struct {
}

func (a *Adidas)makeShoe() IShoe{
	return &AdidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas)makeShirt() IShirt{
	return &AdidasShirt{
		shirt: shirt{
			logo: "adidas",
			size: 14,
		},
	}
}

type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 14,
		},
	}
}

type shoe struct {
	logo string
	size int
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) setSize(size int) {
	s.size = size
}

func (s *shoe) getSize() int {
	return s.size
}

type shirt struct {
	logo string
	size int
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *shirt) getLogo() string {
	return s.logo
}

func (s *shirt) setSize(size int) {
	s.size = size
}

func (s *shirt) getSize() int {
	return s.size
}

type AdidasShirt struct {
	shirt
}

type AdidasShoe struct {
	shoe
}

type NikeShirt struct {
	shirt
}

type NikeShoe struct {
	shoe
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}



