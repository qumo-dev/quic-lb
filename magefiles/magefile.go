// Package main is the mage build script for qumo-lb.
package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Test runs all Go tests.
func Test() error {
	return sh.RunV("go", "test", "./...")
}

// Lint runs golangci-lint.
func Lint() error {
	return sh.RunV("golangci-lint", "run", "./...")
}

// Fmt formats Go code.
func Fmt() error {
	return sh.RunV("go", "fmt", "./...")
}

// Vet runs go vet.
func Vet() error {
	return sh.RunV("go", "vet", "./...")
}

// CI runs the full CI pipeline locally.
func CI() error {
	mg.SerialDeps(Test, Lint, Vet)
	return nil
}
