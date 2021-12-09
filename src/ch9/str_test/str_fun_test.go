package strtest

import (
	"strconv"
	"strings"
	"testing"
)

func TestStrFunc(t *testing.T) {
	s := "A.B.C"
	parts := strings.Split(s, ".")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)

	if i, err := strconv.Atoi("123"); err == nil {
		t.Log(10 + i)
	}

}
