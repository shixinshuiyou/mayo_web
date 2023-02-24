package tool

import "testing"

func TestGenModel(t *testing.T) {
	GenModel("mayo_user", "")
}

func TestGenRepository(t *testing.T) {
	GenDao("mayo_user")
}
