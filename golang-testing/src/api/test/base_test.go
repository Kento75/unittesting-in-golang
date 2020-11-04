package test

import (
	"fmt"
	"github.com/Kento75/unittesting-in-golang/golang-testing/src/api/app"
	"github.com/federicoleon/golang-restclient/rest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	fmt.Println("about to start the application")
	go app.StartApp()

	fmt.Println("application started.")
	os.Exit(m.Run())
}
