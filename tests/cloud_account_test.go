package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/galaxy-future/BridgX/cmd/api/request"
	"github.com/galaxy-future/BridgX/pkg/cloud"
	"github.com/stretchr/testify/assert"
)

const (
	_cloudAccount = _v1Api + `cloud_account/`
)

func TestHello(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "hello world!", w.Body.String())

}
func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()
	lr := request.LoginRequest{
		Username: "root",
		Password: "123456",
	}
	json, _ := json.Marshal(lr)
	fmt.Println(string(json))
	req, _ := http.NewRequest("POST", "/user/login", bytes.NewReader(json))
	req.Header.Set("content-type", "application/json")
	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)
}
func TestCreateCloudAccount(t *testing.T) {
	tests := []request.CreateCloudAccountRequest{
		{
			AccountName:   "test_user",
			Provider:      cloud.BaiduCloud,
			AccountKey:    AKGenerator(cloud.BaiduCloud),
			AccountSecret: "",
		},
		{
			AccountName:   "test_user",
			Provider:      cloud.AwsCloud,
			AccountKey:    AKGenerator(cloud.AwsCloud),
			AccountSecret: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Provider, func(t *testing.T) {
			w := httptest.NewRecorder()
			json, _ := json.Marshal(tt)
			fmt.Println(string(json))
			req, _ := http.NewRequest("POST", _cloudAccount+`create`, bytes.NewReader(json))
			req.Header.Set("content-type", "application/json")
			req.Header.Set("Authorization", "Bearer "+_Token)
			r.ServeHTTP(w, req)
			assert.Equal(t, 200, w.Code)
			fmt.Println(w.Body.String())
		})
	}

}
func TestList(t *testing.T) {
	tests := []struct {
		provider    string
		accountName string
	}{
		{
			provider:    cloud.BaiduCloud,
			accountName: "test_account",
		},
		{
			provider:    cloud.AwsCloud,
			accountName: "test_user",
		},
	}
	for _, tt := range tests {
		t.Run(tt.provider, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", _cloudAccount+`list?provider=`+tt.provider+`&accountName=`+tt.accountName, nil)
			req.Header.Set("content-type", "application/json")
			req.Header.Set("Authorization", "Bearer "+_Token)
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
		})
	}

}
func TestDelete(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", _cloudAccount+`delete/8`, nil)
	req.Header.Set("Authorization", "Bear "+_Token)
	req.Header.Set("content-type", "application/json")
	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)
}
