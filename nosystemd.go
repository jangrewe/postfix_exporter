// +build nosystemd
// This file contains stubs to support non-systemd use

package main

import "io"

type Journal struct {
	io.Closer
	Path string
}

func systemdFlags(enable *bool, unit, slice, path *string) {}

func NewJournal(unit, slice, path string) (*Journal, error) {
	return nil, nil
}

func (e *PostfixExporter) CollectLogfileFromJournal() error {
	return nil
}
