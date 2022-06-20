package mediator

import "fmt"

type ITrain interface {
	arrive()        // 入场
	depart()        // 离场
	permitArrival() // 允许入场
}

type Train struct {
	md   mediator
	name string
}

func (p *Train) arrive() {
	if !p.md.canArrive(p) {
		fmt.Println("列车站台有车,需等待 -----", p.name)
		return
	}
	fmt.Println("列车站台无车----", p.name, "到达列车站台")
}

func (p *Train) depart() {
	fmt.Println("列车", p.name, "离开站台")
	p.md.notifyAboutDeparture()
}

func (p *Train) permitArrival() {
	fmt.Println("列车", p.name, "允许到达")
	p.arrive()
}

func (p *Train) SetMediator(md mediator) {
	p.md = md
}

func (p *Train) GetMediator() mediator {
	return p.md
}

type freightTrain struct {
	Train
}

type mediator interface {
	canArrive(ITrain) bool // 判断列车是否入场
	notifyAboutDeparture() // 监听
}

// 具体中介者
type stationManager struct {
	stationStatus bool
	tranQueue     []ITrain
}

func newStationManger() *stationManager {
	return &stationManager{
		stationStatus: true,
	}
}

func (s *stationManager) canArrive(t ITrain) bool {
	if s.stationStatus {
		s.stationStatus = false
		return true
	}
	s.tranQueue = append(s.tranQueue, t)
	return false
}

func (s *stationManager) notifyAboutDeparture() {
	if !s.stationStatus {
		s.stationStatus = true
	}
	if len(s.tranQueue) > 0 {
		firstTrainInQueue := s.tranQueue[0]
		s.tranQueue = s.tranQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}
