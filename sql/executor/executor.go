package executor

import (
	"github.com/narslan/gleam/flow"
	"github.com/narslan/gleam/sql/expression"
)

type Executor interface {
	Exec() *flow.Dataset
	Schema() expression.Schema
}
