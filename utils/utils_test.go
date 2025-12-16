package utils

import (
	"testing"
)

func Test_hostName(t *testing.T) {
	var _hostName string = HostName()
	if _hostName == "" {
		t.Errorf("hostName function failed to return hostname")
	}
}
