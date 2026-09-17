//go:build tools

package main

// Blank import keeps ariga.io/atlas-provider-gorm in go.mod/go.sum for the
// "load GORM models as schema" step used by migration/atlas.hcl, without
// pulling it into the normal build.
import _ "ariga.io/atlas-provider-gorm/gormschema"
