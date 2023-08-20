package constants

import "errors"

var ErrorInPassword = errors.New("wrong password or email")
var ErrorBloquedUser = errors.New("acount bloqued")
