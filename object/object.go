package object

type Type string

const (
	INTEGER_OBJ      Type = "INTEGER"
	BOOLEAN_OBJ      Type = "BOOLEAN"
	NULL_OBJ         Type = "NULL"
	RETURN_VALUE_OBJ Type = "RETURN_VALUE"
	ERROR_OBJ        Type = "ERROR"
	FUNCTION_OBJ     Type = "FUNCTION"
)

type Object interface {
	Type() Type
	Inspect() string
}
