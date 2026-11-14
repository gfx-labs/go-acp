// Package acp implements the in-development v2 of the Agent Client Protocol.
// Its types come from the draft v2 schema in schema/ (upstream release
// schema-v2.0.0-alpha.15). The draft and this API can change in any release.
//
// The connection code is a copy of the v1 package's. The *_gen.go files come
// from internal/acpgen. Regenerate them with "go generate ." after changing
// the schema files or the generator.
package acp

//go:generate go run -C ../internal/acpgen . -schema ../../v2/schema -out ../../v2
