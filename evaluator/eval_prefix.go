package evaluator

import "monkey/object"

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)

	default:
		return NULL
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE

	case FALSE:
		return TRUE

	case NULL:
		return TRUE

	default:
		return FALSE
	}
}
