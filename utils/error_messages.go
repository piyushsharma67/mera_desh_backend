package utils

import "errors"

var USER_ALREADY_EXISTS = errors.New("user already exists in DB")
var INTERNAL_SERVER_ERROR = errors.New("Internal server error")
