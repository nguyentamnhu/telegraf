package syslog

import (
	"errors"
	"strings"
)

// Framing represents the framing technique we expect the messages to come.
type Framing int

const (
	// OctetCounting indicates the transparent framing technique for syslog transport.
	OctetCounting Framing = iota
	// NonTransparent indicates the non-transparent framing technique for syslog transport.
	NonTransparent
)

func (f Framing) String() string {
	switch f {
	case OctetCounting:
		return "OCTET-COUNTING"
	case NonTransparent:
		return "NON-TRANSPARENT"
	}
	return ""
}

// UnmarshalTOML implements ability to unmarshal framing from TOML files.
func (f *Framing) UnmarshalTOML(data []byte) error {
	return f.UnmarshalText(data)
}

// UnmarshalText implements encoding.TextUnmarshaler
func (f *Framing) UnmarshalText(data []byte) error {
	s := string(data)
	switch strings.ToUpper(s) {
	case `OCTET-COUNTING`:
		fallthrough
	case `"OCTET-COUNTING"`:
		fallthrough
	case `'OCTET-COUNTING'`:
		*f = OctetCounting
		return nil
	case `NON-TRANSPARENT`:
		fallthrough
	case `"NON-TRANSPARENT"`:
		fallthrough
	case `'NON-TRANSPARENT'`:
		*f = NonTransparent
		return nil
	}
	*f = -1
	return errors.New("unknown framing")
}

// MarshalText implements encoding.TextMarshaller
func (f Framing) MarshalText() ([]byte, error) {
	s := f.String()
	if s != "" {
		return []byte(s), nil
	}
	return nil, errors.New("unknown framing")
}
