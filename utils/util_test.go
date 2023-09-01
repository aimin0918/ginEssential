package utils

import "testing"

func TestPhoneDecrypt(t *testing.T) {
	phone := []string{
		"2UxntVYCmeglENEUw8z4lg==",
	}

	for _, v := range phone {
		a := PhoneDecrypt(v)
		t.Log(v, a)
	}
}
