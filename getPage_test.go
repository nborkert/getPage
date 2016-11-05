package getPage

import "testing"

func TestGetContent(t *testing.T) {
	page := "http://www.example.com"
	if x := GetContent(page); x == "" {
		t.Errorf("Did not receive content from ", page)
	}
}

func TestGetSecureContent(t *testing.T) {
	page := "https://example.com:443"
	if x := GetContent(page); x == "" {
		t.Errorf("Did not receive content from ", page)
	}
}
