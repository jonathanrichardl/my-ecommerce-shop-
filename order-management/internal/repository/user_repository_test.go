package repository

import (
	"database/sql"
	"my-ecommerce-shop/order-management/internal/repository/model"
	"my-ecommerce-shop/order-management/internal/testutil"
	"my-ecommerce-shop/order-management/internal/testutil/testdata"
	"my-ecommerce-shop/order-management/internal/util"
	"reflect"
	"testing"
)

var db *sql.DB

func TestMain(m *testing.M) {
	setUpDb()
	m.Run()
	tearDownDB()
}

func TestUserRepository_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{"Should create user succesfully", fields{db}, args{user: testdata.CreateTestUser("userId2", "jorich2", "jonathan12345@jonathan.com", "0123443445")},
			"userId2", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepositoryImplementation{
				db: tt.fields.db,
			}
			got, err := r.Create(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_Get(t *testing.T) {
	expectedUser := testdata.CreateTestUser("userId1", "jorich1", "jonathan123457@jonathan.com",
		"01234434457")

	type fields struct {
		db *sql.DB
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{"should get user succesfully", fields{db}, args{"userId1"}, &expectedUser, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepositoryImplementation{
				db: tt.fields.db,
			}
			got, err := r.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func setUpDb() {
	db = testutil.NewTestDb()
	tearDownDB()

	stmt, err := db.Prepare(`INSERT INTO users 
    (id, username, email, password, phone_number, role) VALUES ($1,$2,$3,$4,$5,$6)`)

	stmt.Exec("userId1", "jorich1", "jonathan123457@jonathan.com", util.EncryptPassword("7772231"),
		"01234434457", model.BUYER)

	if err != nil {
		return
	}
}

func tearDownDB() {
	db.Exec(`TRUNCATE users RESTART IDENTITY;`)
}
