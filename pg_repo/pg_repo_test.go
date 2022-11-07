package pg_repo

import (
	"database/sql"
	"demo/core"
	"demo/utility"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"reflect"
	"testing"
)

var db = utility.InitTestDb()

func TestPgRepository_GetAllCollezioniREPO(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   []core.Collection
	}{
		{
			name:   "Collezione 1",
			fields: fields{db},
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
			r := &PgRepository{
				db: tt.fields.db,
			}
			if got := r.GetAllCollezioniREPO(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCollezioniREPO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgRepository_GetArticoloREPO(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		idArticolo int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   core.Item
	}{
		{
			name:   "Articolo 1",
			fields: fields{db},
			args:   args{1},
			want: core.Item{
				Id:   1,
				Name: "Scarpe",
				Sku:  11,
			},
		}, {
			name:   "Articolo 2",
			fields: fields{db},
			args:   args{2},
			want: core.Item{
				Id:   2,
				Name: "Maglia",
				Sku:  22,
			},
		},
		{
			name:   "Articolo 3",
			fields: fields{db},
			args:   args{3},
			want: core.Item{
				Id:   3,
				Name: "Panta",
				Sku:  33,
			},
		},
		{
			name:   "Articolo 4",
			fields: fields{db},
			args:   args{4},
			want: core.Item{
				Id:   4,
				Name: "Maglione",
				Sku:  44,
			},
		},
		{
			name:   "Articolo 5",
			fields: fields{db},
			args:   args{5},
			want: core.Item{
				Id:   5,
				Name: "Berretto",
				Sku:  55,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PgRepository{
				db: tt.fields.db,
			}
			if got := r.GetArticoloREPO(tt.args.idArticolo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArticoloREPO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgRepository_GetArticoliCollezioneByIdREPO(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		idCollezione int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []core.Item
		wantErr bool
	}{
		{
			name:   "Articoli Collezione 1",
			fields: fields{db: db},
			args:   args{1},
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
			wantErr: false,
		},
		{
			name:   "Collezione 2",
			fields: fields{db: db},
			args:   args{2},
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
			wantErr: false,
		},
		{
			name:    "Collezione vuota",
			fields:  fields{db: db},
			args:    args{0},
			want:    make([]core.Item, 0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PgRepository{
				db: tt.fields.db,
			}
			if got := r.GetArticoliCollezioneByIdREPO(tt.args.idCollezione); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArticoloREPOById() = %v, want %v", got, tt.want)
			}
		})
	}
}
