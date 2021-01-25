package data

import (
	"reflect"
	"testing"

	"github.com/Ariesfall/simple-auth-go/pkg/conn"
	"github.com/jmoiron/sqlx"
)

var (
	db         = conn.Connectdb(conn.SampleConn)
	sampleData = &Tasks{
		ID:         1,
		Name:       "demo",
		DueDate:    "2021-10-01",
		CompltDate: "2021-09-01",
		Status:     0,
	}
)

func TestTaskGet(t *testing.T) {
	type args struct {
		db *sqlx.DB
		in *Tasks
	}
	tests := []struct {
		name    string
		args    args
		want    *Tasks
		wantErr bool
	}{
		{"test1", args{db, sampleData}, sampleData, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TaskGet(tt.args.db, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskCreate(t *testing.T) {
	type args struct {
		db *sqlx.DB
		in *Tasks
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
			if err := TaskCreate(tt.args.db, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("TaskCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTaskUpdate(t *testing.T) {
	type args struct {
		db *sqlx.DB
		in *Tasks
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
			if err := TaskUpdate(tt.args.db, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("TaskUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTaskDelete(t *testing.T) {
	type args struct {
		db *sqlx.DB
		in *Tasks
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
			if err := TaskDelete(tt.args.db, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("TaskDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
