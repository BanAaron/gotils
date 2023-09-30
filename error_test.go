package gotils_test

import (
	"github.com/banaaron/gotils"
	"testing"
)

func TestHandleErrorNil(t *testing.T) {
	var err error
	gotils.HandleError(err)
}
