package gotils_test

import (
	"github.com/banaaron/gotils"
	"testing"
)

func TestHandleErrorNil(t *testing.T) {
	var err error
	res := gotils.HandleError(err)
	if res != nil {
		t.Errorf("handle error nil failed: %s", res)
	}
}
