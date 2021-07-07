package helpers

import (
	"testing"
	"time"

	"github.com/niemeyer/pretty"
	"github.com/stretchr/testify/mock"
)

// Helper functions to check and remove string from a slice of strings.
func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

func AssertCalledEventually(t *testing.T, m *MockAPI, method string, d time.Duration, args ...interface{}) {
	done := make(chan struct{})

	timeout := time.After(d)

	go func() {
		for {
			called := false

			for _, call := range m.Calls {
				if call.Method == method {
					_, differences := mock.Arguments(args).Diff(call.Arguments)

					if differences == 0 {
						// found the expected call
						called = true
						break
					}
				}
			}

			if called {
				done <- struct{}{}
				break
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	select {
	case <-done:
		return
	case <-timeout:
		t.Errorf("Expected method %s to be called with arguments %s", method, pretty.Sprint(args))
		t.Fail()
	}
}
