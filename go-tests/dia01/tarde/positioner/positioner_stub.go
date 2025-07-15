package positioner

type PositionerStub struct {
	FixedDistance float64
}

func (s *PositionerStub) GetLinearDistance(from, to *Position) (linearDistance float64) {
	return s.FixedDistance
}
