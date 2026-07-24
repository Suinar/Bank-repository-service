package operation

import (
	stdlog "log"
	"time"
)

// Started logs an operation lifecycle and returns its completion callback.
func Started(component, layer, name string) func() {
	startedAt := time.Now()
	stdlog.Printf("%s %s: operation started operation=%s", component, layer, name)
	return func() {
		stdlog.Printf(
			"%s %s: operation completed operation=%s duration=%s",
			component, layer, name, time.Since(startedAt),
		)
	}
}

// Failed records a dependency failure without logging request payloads.
func Failed(component, layer, name, dependency string, err error) {
	stdlog.Printf(
		"%s %s: operation failed operation=%s dependency=%s error=%q",
		component, layer, name, dependency, err,
	)
}

// ValidationFailed records a safe validation reason.
func ValidationFailed(component, layer, name, reason string) {
	stdlog.Printf(
		"%s %s: validation failed operation=%s reason=%s",
		component, layer, name, reason,
	)
}

// NilInput records a rejected nil mapping input.
func NilInput(component, layer, name string) {
	stdlog.Printf("%s %s: nil input operation=%s", component, layer, name)
}

// Result records only aggregate result size, never result contents.
func Result(component, layer, name string, count int) {
	stdlog.Printf("%s %s: result operation=%s count=%d", component, layer, name, count)
}
