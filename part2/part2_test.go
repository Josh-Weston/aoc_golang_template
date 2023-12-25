package part2

import (
	"bytes"
	"testing"
)

func TestPart2(t *testing.T) {
	input := `something
here`

	v, err := Run(bytes.NewBufferString(input))
	if err != nil {
		t.Fatal(err)
	}

	expected := -1
	if v != expected {
		t.Fatalf("wrong value; want=%d, got=%d", expected, v)
	}
}
