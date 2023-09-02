package constants

import "errors"

var ErrorInPassword = errors.New("wrong password or email")
var ErrorBloquedUser = errors.New("acount bloqued")
var ErrorProductNotExist = errors.New("product not exist")
var ErrorProductNotProductID = errors.New("error in product id param")
