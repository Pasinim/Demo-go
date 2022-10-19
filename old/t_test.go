//package main
//
//import (
//	"demo/core"
//	"github.com/stretchr/testify/mock"
//	_ "github.com/stretchr/testify/mock"
//	"testing"
//)
//
///**
//1. Genero mock
//2. Metodi che restituiscono
//*/
//
//import (
//	_ "github.com/gin-gonic/gin"
//	_ "github.com/lib/pq"
//)
//
///** I test devono testare ogni livello in modo indipendente -> - Dependency Injection */
//
//// creo una struttura mock che implementa i metodi dell'interfaccia
//type RepositoryMock struct {
//	mock.Mock //marca la struttura come mock
//}
//
//func (m *RepositoryMock) GetArticoloREPO(i int) core.Item {
//	args := m.Called(i)
//	return args.Get(0).(core.Item)
//	return core.Item{
//		Id:   1,
//		Name: "boh",
//		Sku:  "111",
//	}
//}
//
//func TestDoGetArticoliREPO(t *testing.T) {
//	testObj := new(RepositoryMock)
//	testObj.On("GetArticoloREPO", 1).Return(core.Item{
//		Id:   1,
//		Name: "boh",
//		Sku:  "111",
//	})
//
//	// call the code we are testing
//	RepositoryMock.GetArticoloREPO(1)
//	// assert that the expectations were met
//	testObj.AssertExpectations(t)
//}
//func (m *RepositoryMock) GetAllCollezioniREPO() []core.Collection {
//	return nil
//}
//
//func (m *RepositoryMock) GetArticoliREPO() []core.Item {
//	//TODO implement me\
//	panic("implement me")
//}
//
//func (m *RepositoryMock) GetArticoliCollezioneREPO(i int) []core.Item {
//	//TODO implement me
//	panic("implement me")
//}
//
////func TestRepo_GetAllCollezioniREPO(t *testing.T) {
////	var m mock
////	router := api.New(m)
////	w := httptest.NewRecorder()
////	req, err := http.NewRequest("GET", "/collezione/", nil)
////	if err != nil {
////		log.Fatal(err)
////	}
////
////	router.ServeHTTP(w, req)
////	assert.Equal(t, 200, w.Code)
////
////}
