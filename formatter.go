package srslog

import (
	"fmt"
	"os"
	"time"
)

const appNameMaxLength = 48 // limit to 48 chars as per RFC5424

// Formatter is a type of function that takes the consituent parts of a
// syslog message and returns a formatted string. A different Formatter is
// defined for each different syslog protocol we support.
type Formatter func(p Priority, hostname, tag, content string) string

// DefaultFormatter is the original format supported by the Go syslog package,
// and is a non-compliant amalgamation of 3164 and 5424 that is intended to
// maximize compatibility.
func DefaultFormatter(p Priority, hostname, tag, content string) string {
	timestamp := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf("<%d> %s %s %s[%d]: %s",
		p, timestamp, hostname, tag, os.Getpid(), content)
	return msg
}

// UnixFormatter omits the hostname, because it is only used locally.
func UnixFormatter(p Priority, hostname, tag, content string) string {
	timestamp := time.Now().Format(time.Stamp)
	msg := fmt.Sprintf("<%d>%s %s[%d]: %s",
		p, timestamp, tag, os.Getpid(), content)
	return msg
}

// RFC3164Formatter provides an RFC 3164 compliant message.
func RFC3164Formatter(p Priority, hostname, tag, content string) string {
	timestamp := time.Now().Format(time.Stamp)
	msg := fmt.Sprintf("<%d>%s %s %s[%d]: %s",
		p, timestamp, hostname, tag, os.Getpid(), content)
	return msg
}

// if string's length is greater than max, then use the last part
func truncateStartStr(s string, max int) string {
	if len(s) > max {
		return s[len(s)-max:]
	}
	return s
}

// RFC5424Formatter provides an RFC 5424 compliant message.
func RFC5424Formatter(p Priority, hostname, tag, content string) string {
	timestamp := time.Now().Format(time.RFC3339)
	pid := os.Getpid()
	appName := truncateStartStr(os.Args[0], appNameMaxLength)
	msg := fmt.Sprintf("<%d>%d %s %s %s %d %s - %s",
		p, 1, timestamp, hostname, appName, pid, tag, content)
	return msg
}

// RFC5424FormatterWithAppNameAsTag rsyslog uses appname part of syslog message to fill in an %syslogtag% template
// attribute in rsyslog.conf. In order to be backward compatible to rfc3164
// tag will be also used as an appname
func RFC5424FormatterWithAppNameAsTag(p Priority, hostname, tag, content string) string {
	timestamp := time.Now().Format(time.RFC3339)
	pid := os.Getpid()
	msg := fmt.Sprintf("<%d>%d %s %s %s %d %s - %s",
		p, 1, timestamp, hostname, tag, pid, tag, content)
	return msg
}

// RFC5424MicroFormatterWithAppNameAsTag The timestamp field in rfc5424 is derived from rfc3339. Whereas rfc3339 makes allowances
// for multiple syntaxes, there are further restrictions in rfc5424, i.e., the maximum
// resolution is limited to "TIME-SECFRAC" which is 6 (microsecond resolution)
func RFC5424MicroFormatterWithAppNameAsTag(p Priority, hostname, tag, content string) string {
	timestamp := time.Now().Format("2006-01-02T15:04:05.000000Z07:00")
	pid := os.Getpid()
	msg := fmt.Sprintf("<%d>%d %s %s %s %d %s - %s",
		p, 1, timestamp, hostname, tag, pid, tag, content)
	return msg
}
