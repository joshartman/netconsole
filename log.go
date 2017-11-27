package netconsole

import (
	"fmt"
	"regexp"
	"time"
)

// A Log is a log line generated by the netconsole kernel module.
type Log struct {
	// Elapsed is the amount of time elapsed between when a client machine was
	// booted and when this log was sent.
	Elapsed time.Duration

	// Message is the log message.
	Message string
}

// TODO(mdlayher): consider replacing regex parsing.

// logRe matches logs generated by the netconsole kernel module.
var logRe = regexp.MustCompile(`\[\s*(\d+).(\d+)\]\s(.+)`)

// ParseLog parses a log line in the netconsole format.
func ParseLog(s string) (*Log, error) {
	groups := logRe.FindStringSubmatch(s)
	if len(groups) != 4 {
		return nil, fmt.Errorf("malformed netconsole log: %q", s)
	}

	// If the regex matches:
	//   - group 0 is the entire string
	//   - group 1 is the whole number portion of the elapsed duration
	//   - group 2 is the decimal portion of the elapsed duration
	//   - group 3 is the message body

	format := fmt.Sprintf("%s.%ss", groups[1], groups[2])
	elapsed, err := time.ParseDuration(format)
	if err != nil {
		return nil, err
	}

	l := &Log{
		Elapsed: elapsed,
		Message: groups[3],
	}

	return l, nil
}
