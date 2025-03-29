package rfc5322

import "errors"

var ErrorInvalidAddress = errors.New("invalid email address format")
var ErrorInvalidName = errors.New("invalid name format")
var ErrorNeedSender = errors.New("need sender address")
var ErrorNeedToCcBcc = errors.New("need to, cc, or bcc address")
var ErrorInvalidMessageID = errors.New("invalid Message-ID format")
