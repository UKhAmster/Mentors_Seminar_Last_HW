package main

import (
	"strings"
	"testing"
)

func TestWriteTable(t *testing.T) {
	var sb strings.Builder
	writeTable(&sb, 7)

	lines := strings.Split(strings.TrimSpace(sb.String()), "\n")
	if len(lines) != 10 {
		t.Fatalf("ожидалось 10 строк, получено %d", len(lines))
	}
	if lines[0] != "7 x  1 = 7" {
		t.Errorf("первая строка = %q", lines[0])
	}
	if lines[9] != "7 x 10 = 70" {
		t.Errorf("последняя строка = %q", lines[9])
	}
}
