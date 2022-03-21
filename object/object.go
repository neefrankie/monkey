package object

type Type string

const (
	INTEGER_OBJ Type = "INTEGER"
	BOOLEAN_OBJ Type = "BOOLEAN"
	NULL_OBJ    Type = "NULL"
)

type Object interface {
	Type() Type
	Inspect() string
}
