package api

import (
	"demo/core"
	"demo/mockt"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http/httptest"
	"testing"
)

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
			fmt.Println("GOT", got)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEcommerceApi_GETAllCollezioni(t *testing.T) {

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
			name: "AllCollezioni",
			fields: fields{
				rep: mk,
			},
			args: args{
				c: ctx,
			},
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
			e.GETAllCollezioni(tt.args.c)
			var got []core.Collection

			err := json.Unmarshal(w.Body.Bytes(), &got)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEcommerceApi_GETArticoliCollezioneById(t *testing.T) {
	mk := new(mockt.RepositoryMock)

	type fields struct {
		rep Repository
	}
	type args struct {
		r httptest.ResponseRecorder
		c *gin.Context
	}

	tests := []struct {
		name   string
		fields fields
		args   func() args
		want   []core.Item
	}{
		{
			name:   "Collezione senza id",
			fields: fields{mk},
			args: func() args {
				r := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(r)
				return args{
					r: *r,
					c: ctx,
				}
			},
			want: []core.Item(nil),
		},
		{
			name: "Collezione 1",
			fields: fields{
				rep: mk,
			},

			args: func() args {
				r := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(r)
				ctx.AddParam("id", "1")
				return args{
					r: *r,
					c: ctx,
				}
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
					Sku:  33,
				},
			},
		},
		{
			name:   "Collezione 2",
			fields: fields{mk},
			args: func() args {
				r := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(r)
				ctx.AddParam("id", "2")
				return args{
					r: *r,
					c: ctx,
				}

			},
			want: []core.Item{
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

	for idx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EcommerceApi{
				rep: tt.fields.rep,
			}

			mk.On("GetArticoliCollezioneByIdREPO", idx).Return(tt.want)
			a := tt.args()
			e.GETArticoliCollezioneById(a.c)
			var got []core.Item
			res := a.r.Result()

			str, err := io.ReadAll(res.Body)

			if err != nil {
				log.Fatal()
			}

			fmt.Println("BODYdad....", string(str))
			err = json.Unmarshal(a.r.Body.Bytes(), &got)

			fmt.Println("....", got)
			assert.Nil(t, err)
			assert.Equal(t, 200, res.StatusCode)
			assert.Equal(t, tt.want, got)
		})
	}
}

/*
*
GetArticoliBycollezioneid
*/
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
		want   core.Collection
	}{
		{
			name: "Collezione 1",
			fields: fields{
				mk,
			},
			args: args{
				c: ctx,
				//c: context.WithValue(ctx, "id", 1),
			},
			want: core.Collection{
				Id:       1,
				Name:     "Collezione Estiva",
				Articles: nil,
			},
		},
		{
			name:   "Collezione 2",
			fields: fields{mk},
			args: args{
				c: ctx,
				//context.WithValue(ctx, "id", 2),
			},
			want: core.Collection{
				Id:       2,
				Name:     "Collezione Invernale",
				Articles: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			e := &EcommerceApi{
				rep: tt.fields.rep,
			}
			e.GETCollezione(tt.args.c)

		})
	}
}
