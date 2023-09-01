package log

import (
	"strings"
	"testing"
)

type Req struct {
	ShortBytes []byte
	LongBytes  []byte
	Name       string
	Password   string
}

func Test_desensitization(t *testing.T) {
	req := &Req{
		ShortBytes: []byte("1234567890"),
		LongBytes:  []byte("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"),
		Name:       "test",
		Password:   "123456",
	}
	ret := desensitization(req)
	if ret == nil {
		t.Fatal()
	}

	req2 := ret.(*Req)
	if len(req2.ShortBytes) != 10 {
		t.Fatal("[]byte failed")
	}
	if len(req2.LongBytes) > MAX_TRIM_BYTES {
		t.Fatal("[]byte failed")
	}
	if strings.Trim(req2.Password, "*") != "" {
		t.Fatal("Password failed")
	}
}
