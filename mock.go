package main

import (
	"demo/core"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
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
