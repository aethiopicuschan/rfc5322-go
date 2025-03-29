package rfc5322

var encode = true

func EnableEncode() {
	encode = true
}

func DisableEncode() {
	encode = false
}
