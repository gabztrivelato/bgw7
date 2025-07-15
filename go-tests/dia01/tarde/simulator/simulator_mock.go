package simulator

type MockCatchSimulator struct {
	CanCacthResponse bool
	CanCacthCalled   bool
}

func (m *MockCatchSimulator) CanCatch(hunter, prey *Subject) bool {
	m.CanCacthCalled = true
	return m.CanCacthCalled
}
