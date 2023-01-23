package tests

import (
	"bytes"
	// "encoding/json"
	"github.com/gin-gonic/gin"
	// "github.com/rs/xid"
	"apiGorm/controllers"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFichaClienteModel(t *testing.T) {
	t.Parallel()
	r := gin.Default()
	r.GET("/fichaClienteModel", controllers.GetFichaClienteModel)
	req, _ := http.NewRequest("GET", "/fichaClienteModel", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	t.Logf("%#v", len(responseData))
	// assert.Equal(t, responseData, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
	// t.Parallel()
}

func TestPieDistributor(t *testing.T) {

	t.Parallel()
	r := gin.Default()
	r.POST("/grafica", controllers.PieDistributor)
	var jsonReq = []byte(`{"distribuidor": "OTRO", "fecha": "2022"}`)
	buf := bytes.NewReader(jsonReq)
	// err := json.NewEncoder(buf).Encode(jsonReq)
	// if err != nil {
	// 	t.Log(err)
	// }
	req, _ := http.NewRequest("POST", "/grafica", buf)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// t.Logf("%#v", r)
	responseData, _ := ioutil.ReadAll(w.Body)
	t.Logf("%#v", len(responseData))
	assert.Positive(t, len(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
