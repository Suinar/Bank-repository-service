package operation

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOperationLoggingErrors(t *testing.T) {
	output := captureLogs(t)

	Failed("user", "service", "get", "postgres", errors.New("connection refused"))
	ValidationFailed("user", "gRPC handler", "get", "request is nil")
	NilInput("user", "mapper", "user_to_proto")

	require.Contains(t, output.String(), "dependency=postgres")
	require.Contains(t, output.String(), `error="connection refused"`)
	require.Contains(t, output.String(), "validation failed")
	require.Contains(t, output.String(), "nil input")
}
