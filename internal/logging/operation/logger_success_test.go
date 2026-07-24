package operation

import (
	"bytes"
	stdlog "log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOperationLoggingSuccess(t *testing.T) {
	output := captureLogs(t)

	completed := Started("user", "service", "get_all")
	Result("user", "service", "get_all", 2)
	completed()

	require.Contains(t, output.String(), "operation started operation=get_all")
	require.Contains(t, output.String(), "result operation=get_all count=2")
	require.Contains(t, output.String(), "operation completed operation=get_all")
	require.Contains(t, output.String(), "duration=")
}

func captureLogs(t *testing.T) *bytes.Buffer {
	t.Helper()
	previousOutput, previousFlags, previousPrefix := stdlog.Writer(), stdlog.Flags(), stdlog.Prefix()
	output := &bytes.Buffer{}
	stdlog.SetOutput(output)
	stdlog.SetFlags(0)
	stdlog.SetPrefix("")
	t.Cleanup(func() {
		stdlog.SetOutput(previousOutput)
		stdlog.SetFlags(previousFlags)
		stdlog.SetPrefix(previousPrefix)
	})
	return output
}
