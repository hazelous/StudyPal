package health

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_OK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	Handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", res.StatusCode)
	}
	body, _ := io.ReadAll(res.Body)
	if string(body) != "ok" {
		t.Fatalf(`expected body "ok", got %q`, string(body))
	}
}
