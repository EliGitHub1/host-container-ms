package tests

import (
	"host-container-ms/dataBase"
	"testing"
)

func TestDataSourceIsOpen(t *testing.T) {
	if got := dataBase.GetDataBase(); got == nil {
		t.Errorf("fail to connect to db")
	}
}
