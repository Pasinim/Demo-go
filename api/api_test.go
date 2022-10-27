package api

import (
	"demo/core"
	"demo/mock"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestEcommerceApi_GETCollezione(t *testing.T) {
	mk := new(mockt.RepositoryMock)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	type fields struct {
		rep Repository
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []core.Collection
	}{
		{
			name: "Collezione Estiva",
			fields: fields{
				rep: mk,
			},
			args: args{ctx},
			want: []core.Collection{
				{
					Id:       1,
					Name:     "Collezione Estiva",
					Articles: nil,
				},
				{
					Id:       2,
					Name:     "Collezione Invernale",
					Articles: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EcommerceApi{
				rep: tt.fields.rep,
			}

			mk.On("GetAllCollezioniREPO").Return(tt.want)
			e.GETCollezione(tt.args.c)

			var got []core.Collection
			err := json.Unmarshal(w.Body.Bytes(), &got)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

//func TestEcommerceApi_GETArticoli(t *testing.T) {
//	type fields struct {
//		rep Repository
//	}
//	type args struct {
//		c *gin.Context
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			e := &EcommerceApi{
//				rep: tt.fields.rep,
//			}
//			e.GETArticoli(tt.args.c)
//		})
//	}
//}

//func TestNew(t *testing.T) {
//	type args struct {
//		r Repository
//	}
//	tests := []struct {
//		name string
//		args args
//		want *gin.Engine
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := New(tt.args.r); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("New() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
