package yaml_test

import (
	"testing"

	. "github.com/essentialkaos/check"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})
