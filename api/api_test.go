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
			name: "Collezione 1",
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
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEcommerceApi_GETArticoli(t *testing.T) {
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
		want   []core.Item
	}{
		{
			name: "Articoli 1",
			fields: fields{
				rep: mk,
			},
			args: args{
				c: ctx,
			},
			want: []core.Item{

				{
					Id:   1,
					Name: "Scarpe",
					Sku:  11,
				},
				{
					Id:   2,
					Name: "Maglia",
					Sku:  22,
				},
				{
					Id:   3,
					Name: "Panta",
					Sku:  44,
				},
				{
					Id:   4,
					Name: "Maglione",
					Sku:  44,
				},
				{
					Id:   5,
					Name: "Berretto",
					Sku:  55,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EcommerceApi{
				rep: tt.fields.rep,
			}
			mk.On("GetArticoliREPO").Return(tt.want)
			e.GETArticoli(tt.args.c)
			var got []core.Item
			err := json.Unmarshal(w.Body.Bytes(), &got)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
