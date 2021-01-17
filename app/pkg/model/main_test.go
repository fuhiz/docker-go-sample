package model_test

import (
	"testing"

	"github.com/fuhiz/docker-go-sample/app/pkg/connecter"
)

func TestMain(m *testing.M) {
	connecter.Setup()
	m.Run()
}
