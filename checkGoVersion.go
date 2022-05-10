package notify

import (
	"runtime"

	"github.com/hashicorp/go-version"
)

// Make sure go version >= 1.18.2
// Due to the bug in https://github.com/golang/go/issues/52612

func errGoVersion() {
	panic("go-notify: go version must >= 1.18.2. Due to https://github.com/golang/go/issues/52612")
}

func init() {
	v := runtime.Version()
	if v[:2] != "go" {
		errGoVersion()
	}

	v1, err := version.NewVersion(v[2:])
	if err != nil {
		errGoVersion()
	}

	v2, err := version.NewVersion("1.18.2")
	if err != nil {
		errGoVersion()
	}

	if v1.LessThan(v2) {
		errGoVersion()
	}
}
