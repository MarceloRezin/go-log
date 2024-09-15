package golog

import (
	"bytes"
	"strings"
	"testing"
)

const EXPECTED_TEXT = "Test"

func TestInit_nullConfig(t *testing.T) {
	_, err := Init(nil)

	if err == nil {
		t.Error("err should be setted")
	}
}

func TestInit(t *testing.T) {
	var testWriter bytes.Buffer

	config := Config{
		&testWriter,
		GetDefaultInfo(),
		GetDefaultWarning(),
		GetDefaultError(),
	}

	logger, err := Init(&config)

	if err != nil {
		t.Error("err not expected")
	}

	logger.Info.Print(EXPECTED_TEXT)

	result := testWriter.String()

	if !strings.Contains(result, INFO_PREFIX) {
		t.Error("info prefix is expected")
	}

	if !strings.Contains(result, EXPECTED_TEXT) {
		t.Error("text is expected")
	}

	testWriter.Reset()

	logger.Warnig.Print(EXPECTED_TEXT)

	result = testWriter.String()

	if !strings.Contains(result, WARNING_PREFIX) {
		t.Error("warning prefix is expected")
	}

	if !strings.Contains(result, EXPECTED_TEXT) {
		t.Error("text is expected")
	}

	testWriter.Reset()

	logger.Error.Print(EXPECTED_TEXT)

	result = testWriter.String()

	if !strings.Contains(result, ERROR_PREFIX) {
		t.Error("error prefix is expected")
	}

	if !strings.Contains(result, EXPECTED_TEXT) {
		t.Error("text is expected")
	}
}
