package source

import (
	"testing"

	"github.com/johnpena/mock/mockgen/internal/tests/generics"
)

func TestAssert(t *testing.T) {
	var x MockEmbeddingIface[int, float64]
	var _ generics.EmbeddingIface[int, float64] = &x
}
