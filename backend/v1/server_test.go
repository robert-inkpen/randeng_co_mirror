package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/robert-inkpen/randeng_co_mirror/backend/config"
)

func TestWebserver(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Error(err)
	}

	s := NewServer(cfg)
	t.Run("Test Prom Metrics Endpoint", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/api/metrics", nil)
		s.ServeHTTP(w, req)
		if w.Result().StatusCode != http.StatusOK {
			t.Errorf("Expected HTTP status 200, got %v", w.Result().StatusCode)
		}
	})
}
