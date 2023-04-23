package stringbuilder_test

import (
	"strings"
	"testing"

	"github.com/Nigel2392/go-datastructures/stringbuilder"
)

func TestStringBuilder(t *testing.T) {
	var sb = stringbuilder.StringBuilder{}
	sb.Append("Hello")
	sb.Append(" ")
	sb.Append("World")
	sb.Prepend(" ")
	sb.Prepend("World")
	sb.Prepend(" ")
	sb.Prepend("Hello")

	if sb.String() != "Hello World Hello World" {
		t.Fatal("Expected \"Hello World World Hello\" but got ", sb.String())
	}

	if sb.Len() != len("Hello World Hello World") {
		t.Fatal("Expected length ", len("Hello World Hello World"), " but got ", sb.Len())
	}
}

func TestStringBuilder2(t *testing.T) {
	var sb = stringbuilder.StringBuilder{}
	sb.WriteString("Hello")
	sb.WriteString(" ")
	sb.WriteString("World")
	sb.WriteString(" ")
	sb.WriteString("Hello")
	sb.WriteString(" ")
	sb.WriteString("World")

	var regularSB = strings.Builder{}
	var written, err = sb.WriteTo(&regularSB)
	if err != nil {
		t.Fatal("Error executing stringbuilder.WriteTo: ", err)
	}

	if written != int64(len("Hello World Hello World")) {
		t.Fatal("Expected ", len("Hello World Hello World"), " bytes written but got ", written)
	}
}

func TestReset(t *testing.T) {
	sb := stringbuilder.StringBuilder{}
	sb.Append("Hello")
	sb.Append(" ")
	sb.Append("World")
	sb.Reset()

	if sb.String() != "" {
		t.Fatal("Expected empty string but got ", sb.String())
	}

	if sb.Len() != 0 {
		t.Fatal("Expected length 0 but got ", sb.Len())
	}
}
