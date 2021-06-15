// +build tests

package integration

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	buf, err := ioutil.ReadFile("/tmp/env")
	if os.IsNotExist(err) {
		panic("test does not ran in a container")
	}

	if !bytes.Contains(buf, []byte("docker")) {
		panic("test does not ran in a container")
	}

	os.Exit(m.Run())
}

