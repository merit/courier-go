package courier_test

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/merit/courier-go/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCourier_ExecuteRequest(t *testing.T) {
	t.Parallel()
	want := []byte("Ok")
	server := mockServer(want)
	defer server.Close()
	api := &courier.APIConfiguration{
		AuthToken: "",
	}
	req, err := http.NewRequestWithContext(context.Background(), "GET", server.URL, http.NoBody)
	if err != nil {
		t.Fatal(err)
	}
	got, err := api.ExecuteRequest(req)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(string(want), string(got)))
	}
}

func mockServer(want []byte) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(want)
	}))
	return server
}
