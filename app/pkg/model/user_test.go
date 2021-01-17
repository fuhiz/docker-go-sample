package model_test

import (
	"testing"

	"github.com/fuhiz/docker-go-sample/app/pkg/connecter"
	"github.com/fuhiz/docker-go-sample/app/pkg/model"
)

func TestCreateUser(t *testing.T) {
	newUser := model.NewUser("test_user", 30)
	user, _ := model.CreateUser(connecter.DB(), newUser)

	if user.Name != "test_user" {
		t.Fatal("model.CreateUser Failed")
	}
}
