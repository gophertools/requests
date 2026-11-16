package minitrue_test

import (
	"testing"

	"github.com/gophertools/requests/v2/internal/be"
	"github.com/gophertools/requests/v2/internal/minitrue"
)

func TestCond(t *testing.T) {
	be.Equal(t, 1, minitrue.Cond(true, 1, 2))
	be.Equal(t, 2, minitrue.Cond(false, 1, 2))
}
