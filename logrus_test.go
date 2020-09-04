package logx

import (
	"testing"
)

func Test_Logger(t *testing.T) {
	New().Component("log").Category("test").Info("test log")
}
