package gateway

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"strings"
)

func CustomMatcher(key string) (string, bool) {
	if strings.HasPrefix(key, "X-") {
		return key, true
	}

	return runtime.DefaultHeaderMatcher(key)
}
