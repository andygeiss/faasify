package assert

import (
	"fmt"
	"testing"
)

func That(t *testing.T, desc string, got, expected any) {
	t.Run(desc, func(t *testing.T) {
		if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", expected) {
			t.Fatalf("got '%v' but it's not like expected '%v'", got, expected)
		}
	})
}
