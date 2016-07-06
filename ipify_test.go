package ipify

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"src/github.com/stretchr/testify/assert"
	"testing"
)

func TestIpifi(t *testing.T) {
	server := httptest.NewServer(Ipifi())
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Status error %s", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "127.0.0.1", string(data))

	if string(data) != "127.0.0.1" {
		t.Fatalf("Invalid IP %s", data)
	}
}
