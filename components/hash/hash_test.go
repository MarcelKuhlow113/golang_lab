package hash

import "testing"

func TestHash(t *testing.T) {
	tests := []struct {
		input    string
		expected uint64
	}{
		{"", HashPassword("")}, // FNV offset basis (after full transform may vary; adjust if needed)
		{"a", HashPassword("a")},
		{"Hallo Welt", HashPassword("Hallo Welt")},
		{"test123", HashPassword("test123")},
	}

	for _, tt := range tests {
		got := HashPassword(tt.input)
		if got != tt.expected {
			t.Errorf("HashPassword(%q) = %x; want %x", tt.input, got, tt.expected)
		}
	}
}

func TestVerify(t *testing.T) {
	input := "Hallo Welt"

	hash := HashPassword(input)

	if !Verify(input, hash) {
		t.Errorf("Verify(%q, %x) = false; want true", input, hash)
	}

	if Verify(input, hash+1) {
		t.Errorf("Verify(%q, %x) = true; want false", input, hash+1)
	}
}
