package mocker

import "github.com/stretchr/testify/mock"

type MockCall struct {
	Calls   []*mock.Call
	GenMock func() *mock.Call
}

func NewMockCall(genMockFn func() *mock.Call) *MockCall {
	return &MockCall{
		Calls:   []*mock.Call{},
		GenMock: genMockFn,
	}
}

func (m *MockCall) NumberCallReturn(number int, value ...interface{}) *mock.Call {
	idx := number - 1
	if len(m.Calls) >= number {
		m.Calls[idx].Return(value...)
		return m.Calls[idx]
	}

	genMock := m.GenMock()
	genMock.Return(value...)
	m.Calls = append(m.Calls, genMock)
	return m.Calls[idx]
}
