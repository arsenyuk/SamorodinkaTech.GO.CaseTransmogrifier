package casetransmogrifier

import "testing"

func TestSplitBySpace(t *testing.T) {

	words := SplitBySpace("1 aA")
	if len(words) != 2 {
		t.Errorf("SplitBySpace Expected 2, Received %v", len(words))
	}

	words = SplitBySpace("Empl ID")
	if len(words) != 2 {
		t.Errorf("SplitBySpace Expected 2, Received %v", len(words))
	}
}

func TestSplitByNewline(t *testing.T) {

	words := SplitByNewline("\r")
	if len(words) != 1 {
		t.Errorf("SplitByNewline Expected 1, Received %v", len(words))
	}

	words = SplitByNewline("\n")
	if len(words) != 2 {
		t.Errorf("SplitByNewline Expected 2, Received %v", len(words))
	}

	words = SplitByNewline("\n\r")
	if len(words) != 2 {
		t.Errorf("SplitByNewline Expected 2, Received %v", len(words))
	}

	words = SplitByNewline("\r\n\n")
	if len(words) != 3 {
		t.Errorf("SplitByNewline Expected 3, Received %v", len(words))
	}

	words = SplitByNewline("\r\n\r\n")
	if len(words) != 3 {
		t.Errorf("SplitByNewline Expected 3, Received %v", len(words))
	}

	words = SplitByNewline("\r\n\r\n\r")
	if len(words) != 3 {
		t.Errorf("SplitByNewline Expected 3, Received %v", len(words))
	}

	words = SplitByNewline("\r\n\r\n\n")
	if len(words) != 4 {
		t.Errorf("SplitByNewline Expected 4, Received %v", len(words))
	}
}
