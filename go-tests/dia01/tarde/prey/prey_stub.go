package prey

import (
	"testdoubles/positioner"
)

type PreyStub struct {
	Speed    float64
	Position *positioner.Position
}

func (p *PreyStub) GetSpeed() (speed float64) {
	return p.Speed
}

func (p *PreyStub) GetPosition() positioner.Position {
	return *p.Position
}
