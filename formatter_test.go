package srslog

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestDefaultFormatter(t *testing.T) {
	out := DefaultFormatter(LOG_ERR, "hostname", "tag", "content")
	expected := fmt.Sprintf("<%d> %s %s %s[%d]: %s",
		LOG_ERR, time.Now().Format(time.RFC3339), "hostname", "tag", os.Getpid(), "content")
	if out != expected {
		t.Errorf("expected %v got %v", expected, out)
	}
}

func TestUnixFormatter(t *testing.T) {
	out := UnixFormatter(LOG_ERR, "hostname", "tag", "content")
	expected := fmt.Sprintf("<%d>%s %s[%d]: %s",
		LOG_ERR, time.Now().Format(time.Stamp), "tag", os.Getpid(), "content")
	if out != expected {
		t.Errorf("expected %v got %v", expected, out)
	}
}

func TestRFC3164Formatter(t *testing.T) {
	out := RFC3164Formatter(LOG_ERR, "hostname", "tag", "content")
	expected := fmt.Sprintf("<%d>%s %s %s[%d]: %s",
		LOG_ERR, time.Now().Format(time.Stamp), "hostname", "tag", os.Getpid(), "content")
	if out != expected {
		t.Errorf("expected %v got %v", expected, out)
	}
}

func TestRFC5424Formatter(t *testing.T) {
	out := RFC5424Formatter(LOG_ERR, "hostname", "tag", "content")
	expected := fmt.Sprintf("<%d>%d %s %s %s %d %s - %s",
		LOG_ERR, 1, time.Now().Format(time.RFC3339), "hostname", truncateStartStr(os.Args[0], appNameMaxLength),
		os.Getpid(), "tag", "content")
	if out != expected {
		t.Errorf("expected %v got %v", expected, out)
	}
}

func TestRFC5424FormatterWithAppNameAsTag(t *testing.T) {
	out := RFC5424FormatterWithAppNameAsTag(LOG_ERR, "hostname", "tag", "content")
	expected := fmt.Sprintf("<%d>%d %s %s %s %d %s - %s",
		LOG_ERR, 1, time.Now().Format(time.RFC3339), "hostname", "tag", os.Getpid(), "tag", "content")
	if out != expected {
		t.Errorf("expected %v got %v", expected, out)
	}
}

func TestRFC5424MicroFormatterWithAppNameAsTag(t *testing.T) {
	out := RFC5424MicroFormatterWithAppNameAsTag(LOG_ERR, "hostname", "tag", "content")
	expected := fmt.Sprintf("<%d>%d %s %s %s %d %s - %s",
		LOG_ERR, 1, "%s", "hostname", "tag", os.Getpid(), "tag", "content")
	var timestamp string
	if n, err := fmt.Sscanf(out, expected, &timestamp); err != nil || n != 1 {
		t.Errorf("output = '%q', didn't match '%q' (%d %s)", out, expected, n, err)
	}
}

func TestTruncateStartStr(t *testing.T) {
	out := truncateStartStr("abcde", 3)
	if strings.Compare(out, "cde") != 0 {
		t.Errorf("expected \"cde\" got %v", out)
	}
	out = truncateStartStr("abcde", 5)
	if strings.Compare(out, "abcde") != 0 {
		t.Errorf("expected \"abcde\" got %v", out)
	}
}
