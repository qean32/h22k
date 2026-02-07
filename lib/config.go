package lib

import (
	"main/model"
)

var KEY_FUNCTION = map[string]func(e model.Event){
	"hash":         func(e model.Event) {}, // hash "ключ" "сообщество" "хеш" (-nl)
	"get":          func(e model.Event) {}, // hash "ключ" (-nl)
	"dihash":       func(e model.Event) {},
	"generatehash": func(e model.Event) {},
	"drop":         func(e model.Event) {},
	"clearlog":     func(e model.Event) {},
}

var FLAG = map[string]string{
	"f":     "",
	"force": "",
	"nl":    "",
}

const PROJECT_NAME = "Holo"
