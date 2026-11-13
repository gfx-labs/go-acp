// Package acp implements the Agent Client Protocol: JSON-RPC 2.0 over
// newline-delimited messages, with the protocol's types generated from the
// schema files in schema/.
//
// The *_gen.go files come from internal/acpgen. Regenerate them with
// "go generate ." after changing the schema files or the generator.
package acp

//go:generate go run -C internal/acpgen . -schema ../../schema -out ../..
