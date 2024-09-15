package golog

import (
	"log"
	"testing"
)

const EXPECTED_INFO_PREFIX = "INFO: "
const EXPECTED_INFO_FLAG = log.Ldate | log.Ltime

const EXPECTED_WARNING_PREFIX = "WARNING: "
const EXPECTED_WARNING_FLAG = log.Ldate | log.Ltime | log.Lshortfile

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

func TestGetDefaultWarning(t *testing.T) {
	op := GetDefaultWarning()

	prefixActual := op.Prefix
	if prefixActual != EXPECTED_WARNING_PREFIX {
		t.Log("prefix should be", EXPECTED_WARNING_PREFIX, "but got", prefixActual)
		t.Fail()
	}

	flagActual := op.Flag
	if flagActual != EXPECTED_WARNING_FLAG {
		t.Log("flag should be", EXPECTED_WARNING_FLAG, "but got", flagActual)
		t.Fail()
	}
}
