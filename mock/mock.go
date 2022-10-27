package mockt

import (
	"demo/core"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) GetAllCollezioniREPO() []core.Collection {
	args := m.Called()
	var ret []core.Collection
	value := args.Get(0)
	if value != nil {
		ret = value.([]core.Collection)
	}
	return ret
}

func (m *RepositoryMock) GetArticoloREPO(i int) *core.Item {
	args := m.Called()
	var ret *core.Item
	v := args.Get(0)
	if v != nil {
		ret = v.(*core.Item)
	}
	return ret
}

func (m *RepositoryMock) GetArticoliCollezioneREPO(i int) []core.Item {
	args := m.Called(i)
	var ret []core.Item
	value := args.Get(0)
	if value != nil {
		ret = value.([]core.Item)
	}
	return ret

}

func (m *RepositoryMock) GetArticoliREPO() []core.Item {
	args := m.Called()

	var r []core.Item
	v := args.Get(0)
	if v != nil {
		r = v.([]core.Item)
	}
	return r

}
