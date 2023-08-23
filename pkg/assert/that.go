package assert

import (
	"fmt"
	"testing"
)

func That(desc string, t *testing.T, got, expected any) {
	t.Run(desc, func(t *testing.T) {
		if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", expected) {
			t.Fatalf("got '%v' but it's not like expected '%v'", got, expected)
		}
	})
}
