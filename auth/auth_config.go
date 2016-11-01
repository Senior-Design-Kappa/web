package auth

import (
	"encoding/base64"
)

// TODO: change these?
const (
	cookieStoreKeyBase64  = "ayrxx98d3tW7+8I4/6QAnrziqns+oCcbraEYB+dzFcFCbBlNm4X75CAK/v3rMW55KsCKx9JU7VnjiyjwYk/QlQ=="
	sessionStoreKeyBase64 = "HQgNa7UkPwLwsl2rez5YrhkOdAZXzpjBJbWESiv6hPcjfU60T910g5OnMKRxoYGCtbkNkCqDgBD6YFMbtoYFag=="
	sessionCookieName     = "session_cookie"
	xsrfName              = "csrf_token"
)

var cookieStoreKey, _ = base64.StdEncoding.DecodeString(cookieStoreKeyBase64)
var sessionStoreKey, _ = base64.StdEncoding.DecodeString(sessionStoreKeyBase64)
