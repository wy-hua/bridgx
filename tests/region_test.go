package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/galaxy-future/BridgX/pkg/cloud"
	"github.com/stretchr/testify/assert"
)

func TestListRegions(t *testing.T) {
	tests := []struct {
		provider string
	}{
		{
			provider: cloud.BaiduCloud,
		},
		{
			provider: cloud.AWSCloud,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", _v1Api+"region/list?provider="+tt.provider, nil)
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(7 * time.Second)
		})
	}

}
func TestListZones(t *testing.T) {
	tests := []struct {
		provider string
		regionId string
	}{
		{
			provider: cloud.BaiduCloud,
			regionId: "bj",
		},
		{
			provider: cloud.AWSCloud,
			regionId: "cn-north-1",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", _v1Api+fmt.Sprintf("zone/list?provider=%s&region_id=%s", tt.provider, tt.regionId), nil)
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(7 * time.Second)
		})
	}

}
