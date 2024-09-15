package golog

import (
	"log"
	"testing"
)

const EXPECTED_INFO_PREFIX = "INFO: "
const EXPECTED_INFO_FLAG = log.Ldate | log.Ltime

func TestGetDefaultInfo(t *testing.T) {
	op := GetDefaultInfo()

	prefixActual := op.Prefix
	if prefixActual != EXPECTED_INFO_PREFIX {
		t.Log("prefix should be", EXPECTED_INFO_PREFIX, "but got", prefixActual)
		t.Fail()
	}

	flagActual := op.Flag
	if flagActual != EXPECTED_INFO_FLAG {
		t.Log("flag should be", EXPECTED_INFO_FLAG, "but got", flagActual)
		t.Fail()
	}
}
