package evaluator

import (
	"lorikeet/object"
)

var builtins = map[string]*object.Builtin{
	"len":  object.GetBuiltinByName("len"),
	"say":  object.GetBuiltinByName("say"),
	"head": object.GetBuiltinByName("head"),
	"last": object.GetBuiltinByName("last"),
	"tail": object.GetBuiltinByName("tail"),
	"push": object.GetBuiltinByName("push"),
}
