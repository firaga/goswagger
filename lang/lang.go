package lang

import (
	"goswagger/constants"
)

type Tags map[string]Texts
type Texts struct {
	texts map[string]string
}

var message = map[string]Texts{"en": enTexts}

func Get(msgKey string) string {
	messages := (message)[constants.CURREBTLANG]
	if message, ok := messages.texts[msgKey]; ok {
		return message
	}
	return msgKey
}
