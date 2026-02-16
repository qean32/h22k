package lib

import (
	"fmt"
	"main/constants"
	"main/deep"
	"main/model"
	"os"
)

var KEY_FUNCTION = map[string]func(e model.Event){
	"hash":   func(e model.Event) {}, // hash ключ пароль {сообщение} (-nl)
	"dihash": func(e model.Event) {}, // dihash ключ пароль (-nl)
	"g:key":  func(e model.Event) {}, // g:key
	"c:log": func(e model.Event) {
		deep.ClearFile(constants.LOG_PATH)
	}, // c:log (-nl -f)
	"master": func(e model.Event) {
		deep.CreateFile(constants.LOG_PATH)
		deep.CreateFile("/prime.asc")
	}, // master (-nl)
	"drop": func(e model.Event) { os.RemoveAll(constants.Root); os.Mkdir(constants.Root, 0755) }, // drop (-f)
	"stop": func(e model.Event) { os.Exit(0) },                                                   // stop
	"help": func(e model.Event) { fmt.Println(constants.HelpMessage) },                           // help
}

var KEY_PARSE = map[string]func(arr []string) (event model.Event, _error bool){
	"hash": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{}
		_error = false
		payload := getPaylaod(arr)

		if len(arr) < 3 || payload == "" {
			_error = true
			return
		}

		event = model.Event{
			Time:     deep.NewTime(),
			Key:      arr[0],
			Password: arr[1],
			Payload:  payload,
			Flags:    filterIsFlag(arr),
		}
		return
	},
	"g:key": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time: deep.NewTime(),
			Key:  arr[0],
		}
		_error = false
		return
	},
	"g:log": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time: deep.NewTime(),
			Key:  arr[0],
		}
		_error = false
		return
	},
	"c:log": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time: deep.NewTime(),
			Key:  arr[0],
		}
		_error = false
		return
	},
	"dihash": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{}
		_error = false

		if len(arr) < 3 {
			_error = true
			return
		}

		event = model.Event{
			Time:     deep.NewTime(),
			Key:      arr[0],
			Password: arr[1],
			Flags:    filterIsFlag(arr),
		}
		return
	},
	"master": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time: deep.NewTime(),
			Key:  arr[0],
		}
		_error = false
		return
	},
	"drop": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time: deep.NewTime(),
			Key:  arr[0],
		}
		_error = false
		return
	},
	"stop": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time: deep.NewTime(),
			Key:  arr[0],
		}
		_error = false
		return
	},
	"help": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time: deep.NewTime(),
			Key:  arr[0],
		}
		_error = false
		return
	},
}

var FLAG = map[string]string{
	"f":     "",
	"force": "",
	"nl":    "",
}
