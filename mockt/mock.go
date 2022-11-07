package mockt

import (
	"demo/core"
	"github.com/stretchr/testify/mock"
)

/*
*
Il mockt permette di creare una struttura con gli stessi metodi
della repository, definisco i metodi con la stessa signature
e decido io cosa devono restituire (dico cosa restituire nel test)
*/
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

func (m *RepositoryMock) GetArticoloREPO(i int) core.Item {
	args := m.Called()
	var ret core.Item
	v := args.Get(0)
	if v != nil {
		ret = v.(core.Item)
	}
	return ret
}

func (m *RepositoryMock) GetArticoliCollezioniREPO() []core.Item {
	args := m.Called()
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

func (m *RepositoryMock) GetArticoliCollezioneByIdREPO(i int) []core.Item {
	args := m.Called(i)
	var r []core.Item
	v := args.Get(0)
	if v != nil {
		r = v.([]core.Item)
	}
	return r
}
