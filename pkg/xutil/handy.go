package xutil

// IfElse returns either the value of "ifTrue" or "ifFalse" depending on whether
// the given "condition" value is true or false.
//
// @param condition A boolean condition.
//
// @param ifTrue The value to be returned if the given condition is true.
//
// @param ifFalse The value to be returned if the given condition is false.
//
// @returns Either the value of "ifTrue" or the value of "ifFalse" depending on
// the boolean value of "condition".
func IfElse[V interface{}](condition bool, ifTrue V, ifFalse V) V {
	if condition {
		return ifTrue
	} else {
		return ifFalse
	}
}
