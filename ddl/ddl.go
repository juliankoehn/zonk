package ddl

import (
	"github.com/juliankoehn/zonk/parser"
)

type migration struct {
	Name       string
	Statements []*parser.Statement
}

type logger struct {
	Enabled bool
	Package string
}

type migrationParams struct {
	Package    string
	Dialect    string
	Migrations []migration
	Logger     logger
}
