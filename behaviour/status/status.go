package status

import "fmt"

type Machine struct {
	state *statusMachine
	goods []*Good
	p     *param
}

type param struct {
	goodNo    int
	goodNum   int
	goodValue int
	money     int
}

func (m *Machine) Run() {
	if err := m.state.run(m.p); err != nil {
		panic(err)
	}
}

func NewMachine() *Machine {
	return &Machine{
		state: nil,
		goods: make([]*Good, 0),
		p: &param{
			goodNo: -1,
		},
	}
}

func (m *Machine) GetAllGood() {
	for _, v := range m.goods {
		fmt.Println(*v)
	}
}

func (m *Machine) AddGood(good *Good) {
	m.goods = append(m.goods, good)
}

func (m *Machine) Prepare() {
	m.state = &statusMachine{[]state{&SelectGoodsState{goods: m.goods}, &MoneyState{goods: m.goods}, &settleState{m.goods}}, 0}
}

func (m *Machine) SelectGood(goodNo, num int) {
	m.p.goodNo = goodNo
	m.p.goodNum = num
}

func (m *Machine) GiveMoney(money int) {
	m.p.money = money
}

// 状态机
type statusMachine struct {
	statuses []state
	curState int
}

func (s *statusMachine) run(p *param) error {
	if s.curState > len(s.statuses)-1 {
		return fmt.Errorf("流程已结束")
	}
	if err := s.statuses[s.curState].handle(p); err != nil {
		return err
	}
	if len(s.statuses) > s.curState {
		s.curState++
	}
	return nil
}

type state interface {
	handle(p *param) error
}

type Good struct {
	name  string
	value int
	count int
}

func (g *Good) GetSurplus() int {
	return g.count
}

func (g *Good) OutGood(num int) {
	g.count = g.count - num
}

// SelectGoodsState 商品状态机
type SelectGoodsState struct {
	goods []*Good
}

func (s *SelectGoodsState) handle(p *param) error {
	good, err := s.GetGood(p.goodNo)
	if err != nil {
		return err
	}
	if good.count < p.goodNum {
		return fmt.Errorf("商品编号(%d)库存不足", p.goodNo)
	}
	return nil
}

func (m *SelectGoodsState) GetGood(no int) (*Good, error) {
	if no > len(m.goods)-1 || no < 0 {
		return nil, fmt.Errorf("错误的编号")
	}
	return m.goods[no], nil
}

// MoneyState 资金状态机
type MoneyState struct {
	goods []*Good
}

func (s *MoneyState) handle(p *param) error {
	err := s.MoneyEnough(p)
	if err != nil {
		return err
	}
	return nil
}

func (m *MoneyState) MoneyEnough(p *param) error {
	good, _ := m.GetGood(p.goodNo)
	if p.money < good.value*p.goodNum {
		return fmt.Errorf("付款前不够请加钱")
	}
	return nil
}

func (m *MoneyState) GetGood(no int) (*Good, error) {
	if no > len(m.goods)-1 || no < 0 {
		return nil, fmt.Errorf("错误的编号")
	}
	return m.goods[no], nil
}

// 结算状态机
type settleState struct {
	goods []*Good
}

func (s *settleState) handle(p *param) error {
	err := s.returnGoodAndMoney(p)
	if err != nil {
		return err
	}
	return nil
}

// 返还余额
func (s *settleState) returnGoodAndMoney(p *param) error {
	fmt.Printf("商品%d出库\n", p.goodNo)
	s.goods[p.goodNo].OutGood(p.goodNum)
	fmt.Printf("返还金额%d\n", p.money-s.goods[p.goodNo].value*p.goodNum)
	return nil
}
