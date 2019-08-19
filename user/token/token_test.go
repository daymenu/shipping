package token

import (
	"strings"
	"testing"

	pb "github.com/daymenu/shipping/user/proto/user"
)

var (
	user = &pb.User{
		Id:    1,
		Email: "qwe.163.com",
	}
)

type MockRepo struct{}

func (repo *MockRepo) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	return users, nil
}

func (repo *MockRepo) Page(req *pb.Request) ([]*pb.User, error) {
	var users []*pb.User
	return users, nil
}

func (repo *MockRepo) Get(id int64) (*pb.User, error) {
	return user, nil
}

func (repo *MockRepo) GetByEmail(email string) (*pb.User, error) {
	return user, nil
}

func (repo *MockRepo) Create(user *pb.User) error {
	return nil
}

func (repo *MockRepo) Update(user *pb.User) error {
	return nil
}

func NewInstance() Authable {
	repo := &MockRepo{}
	return &TokenService{repo}
}

func TestCanCreateToken(t *testing.T) {
	srv := NewInstance()
	token, err := srv.Encode(user)
	if err != nil {
		t.Fail()
	}
	if token == "" {
		t.Fail()
	}
	if len(strings.Split(token, ".")) != 3 {
		t.Fail()
	}
}

func TestDecode(t *testing.T) {
	srv := NewInstance()
	token, err := srv.Encode(user)
	if err != nil {
		t.Fail()
	}
	claims, err := srv.Decode(token)
	if err != nil {
		t.Fail()
	}
	if claims.User == nil {
		t.Fail()
	}
	if claims.User.Email != user.Email {
		t.Fail()
	}
}
