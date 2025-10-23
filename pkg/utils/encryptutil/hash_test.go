package encryptutil

import (
	"testing"
	"time"
)

func TestHashPassword(t *testing.T) {
	t.Log(HashPassword("mima123."))
}

func TestVerifyPassword(t *testing.T) {
	t.Log(time.Now().UnixMilli())
	t.Log(VerifyPassword("$2a$10$9jkwijpnwtLU2YEsHSXmkeWDlikHID9o7j7qc7JlNeMMheI4DZGWy", "mima123."))
}
