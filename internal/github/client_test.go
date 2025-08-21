package github

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserEvents(t *testing.T) {
	mockResponse := `[{"type":"PushEvent","repo":{"name":"NERFTHISPLS/test-repo"}}]`
	parsedJson := Event{
		Type: "PushEvent",
		Repo: Repository{
			Name: "NERFTHISPLS/test-repo",
		},
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, mockResponse)
	}))
	defer server.Close()

	client := &Client{
		http:     server.Client(),
		basePath: server.URL,
	}

	events, err := client.UserEvents("NERFTHISPLS")
	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	if parsedJson != events[0] {
		t.Errorf("expected %v, got %v", parsedJson, events[0])
	}
}
