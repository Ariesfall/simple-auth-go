package data

import (
	"reflect"
	"testing"

	"github.com/Ariesfall/simple-auth-go/pkg/conn"
	"github.com/jmoiron/sqlx"
)

var (
	db         = conn.Connectdb(conn.SampleConn)
	sampleData = &Usrs{
		ID:         1,
		Name:       "demo",
		DueDate:    "2021-10-01",
		CompltDate: "2021-09-01",
		Status:     0,
	}
)

func TestUserGet(t *testing.T) {
	type args struct {
		db *sqlx.DB
		in *Usrs
	}
	tests := []struct {
		name    string
		args    args
		want    *Usrs
		wantErr bool
	}{
		{"test1", args{db, sampleData}, sampleData, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UserGet(tt.args.db, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserCreate(t *testing.T) {
	type args struct {
		db *sqlx.DB
		in *Usrs
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{db, sampleData}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UserCreate(tt.args.db, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("UserCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserUpdate(t *testing.T) {
	type args struct {
		db *sqlx.DB
		in *Usrs
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{db, sampleData}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UserUpdate(tt.args.db, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("UserUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDelete(t *testing.T) {
	type args struct {
		db *sqlx.DB
		in *Usrs
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{db, sampleData}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UserDelete(tt.args.db, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("UserDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
