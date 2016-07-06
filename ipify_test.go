package ipify

import (
	"net/http/httptest"
	"testing"
)

func TestIpifi(t *testing.T) {
	server := httptest.NewServer(Ipifi())
	defer server.Close()

	c := Client{URL: server.URL}
	ip, err := c.CheckIP()
	if err != nil {
		t.Fatal(err)
	}

	if ip != "127.0.0.1" {
		t.Fatalf("Invalid IP %s", ip)
	}
}
