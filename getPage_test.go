package getPage

import "testing"

func TestPrintMessage (t *testing.T) {
	if x := PrintMessage(); x == "" {
		t.Errorf("Expected text, got empty string")
	}

}

func TestPrintMessageEcho (t *testing.T) {
	if x := PrintEcho("Hello"); x != "Hello" {
		t.Errorf("Expected Hello, got ", x)
	}

}
