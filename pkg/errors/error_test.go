package errors

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	request := NewBadRequest("invalid_email_format: %s", "sadf@gmail.")
	fmt.Printf("%#v", request)
}
