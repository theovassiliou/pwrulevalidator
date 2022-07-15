package main_test

import (
	"os"
	"testing"

	main "github.com/theovassiliou/pwrulevalidator"
)

var a main.App

func TestMain(m *testing.M) {
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	a = main.App{}
}

func clearTable() {

}
