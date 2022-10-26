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

func TestPgRepository_GetArticoliCollezioneREPO(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		idCollezione int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []core.Item
	}{
		{
			name:   "Articoli Collezione1",
			fields: fields{db},
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PgRepository{
				db: tt.fields.db,
			}
			if got := r.GetArticoliCollezioneREPO(tt.args.idCollezione); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArticoliCollezioneREPO() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
//func TestPgRepository_GetAllCollezioniREPO(t *testing.T) {
//	type fields struct {
//		db *sql.DB
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   []core.Collection
//	}{
//		{
//			name:   "",
//			fields: fields{},
//			want:   nil,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &PgRepository{
//				db: tt.fields.db,
//			}
//			if got := r.GetAllCollezioniREPO(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetAllCollezioniREPO() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestPgRepository_GetArticoliREPO(t *testing.T) {
//	type fields struct {
//		db *sql.DB
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   []core.Item
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &PgRepository{
//				db: tt.fields.db,
//			}
//			if got := r.GetArticoliREPO(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetArticoliREPO() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestPgRepository_GetArticoloREPO(t *testing.T) {
//	type fields struct {
//		db *sql.DB
//	}
//	type args struct {
//		idArticolo int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   core.Item
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &PgRepository{
//				db: tt.fields.db,
//			}
//			if got := r.GetArticoloREPO(tt.args.idArticolo); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetArticoloREPO() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
