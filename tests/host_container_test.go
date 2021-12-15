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

//func TestHostGetAllHosts(t *testing.T) {
//
//	r := Router.SetupRouter()
//	var w = httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/host", nil)
//	r.ServeHTTP(w, req)
//	fmt.Println(req)
//	assert.Equal(t, 200, w.Code)
//assert.Equal(t, "pong", w.Body.String())
//}
