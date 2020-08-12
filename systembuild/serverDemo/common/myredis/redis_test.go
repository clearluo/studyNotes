package myredis

import (
	"serverDemo/common/consts"
	"testing"
)

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Set("key", "value", MINUTE)
	}
}

func TestStudy(t *testing.T) {
	Subscribe(consts.SENSITIVE_AFTER)
}
