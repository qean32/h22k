package lib

import (
	"fmt"
	"main/constants"
	"main/deep"
	"main/model"
	"os"
	"slices"
	"strings"
)

var KEY_FUNCTION = map[string]model.EventFunction{
	"cripto": func(e model.Event) {
		deep.PushToFile(constants.DATA_PATH, fmt.Sprintf("%s %s", e.KeyWord, e.Payload))
		deep.TMP_DATA = append(deep.TMP_DATA, []string{e.KeyWord, e.Payload})
	},
	"ecripto": func(e model.Event) {
		index := slices.IndexFunc(deep.TMP_DATA, func(item []string) bool {
			return item[0] == e.KeyWord
		})

		if index != -1 {
			fmt.Println(deep.TMP_DATA[index][1])
		} else {
			fmt.Println(constants.UNDEFINED_WORD_KEY)
		}
	},
	"g:key": func(e model.Event) {},
	"c:log": func(e model.Event) {
		deep.ClearFile(constants.LOG_PATH)
	},
	"master": func(e model.Event) {
		deep.CreateFile(constants.LOG_PATH)
		deep.CreateFile(constants.COMMAND_PATH)
		deep.CreateFile(constants.DATA_PATH)
	},
	"drop": func(e model.Event) {
		os.RemoveAll(constants.Root)
		os.Mkdir(constants.Root, 0755)
	},
	"stop": func(e model.Event) { os.Exit(0) },
	"help": func(e model.Event) {
		fmt.Println(constants.HelpMessage)
	},
	"place": func(e model.Event) {
		deep.PushToFile(constants.COMMAND_PATH, fmt.Sprintf("%s %s", e.KeyWord, e.Payload))
		deep.TMP_COMMANDS = append(deep.TMP_COMMANDS, []string{e.KeyWord, e.Payload})
	},
	"run": func(e model.Event) {
		index := slices.IndexFunc(deep.TMP_COMMANDS, func(item []string) bool {
			return item[0] == strings.TrimSpace(e.KeyWord)
		})

		if index != -1 {
			deep.RunCommand(deep.TMP_COMMANDS[index][1])
		} else {
			fmt.Println(constants.UNDEFINED_WORD_KEY)
		}
	},
	"run:m": func(e model.Event) {
		index := slices.IndexFunc(deep.TMP_COMMANDS, func(item []string) bool {
			return item[0] == strings.TrimSpace(e.KeyWord)
		})

		if index != -1 {
			tmp := strings.Split(deep.TMP_COMMANDS[index][1], ";")

			for i := 0; i < len(tmp); i++ {
				deep.RunCommand(tmp[i])
			}
		} else {
			fmt.Println(constants.UNDEFINED_WORD_KEY)
		}
	},
	"comm": func(e model.Event) {
		fmt.Println("")
		fmt.Print("- ")
		fmt.Println(strings.Join(deep.ReadFile(constants.COMMAND_PATH), "\n- "))
		fmt.Println("")
	},
	"rm:c": deep.DECORATOR_ACCESS_ACTION(
		func(e model.Event) {
			filtered := FILTER(deep.TMP_COMMANDS, func(item []string) bool { return item[0] != e.KeyWord })
			deep.TMP_COMMANDS = filtered
			deep.WriteFile(strings.Join(deep.MatrixToArrayString(filtered), "\n"), constants.COMMAND_PATH)
		}),
}

var KEY_PARSE = map[string]model.FnRerutnEvent{
	"cripto": func(arr []string) (e model.Event, _error bool) {
		payload := getPaylaod(arr)

		if len(arr) < 3 || payload == "" {
			_error = true
			return
		}

		e = model.Event{
			Time:     deep.NewTime(),
			Key:      arr[0],
			KeyWord:  arr[1],
			Password: arr[2],
			Payload:  payload,
			Flags:    filterIsFlag(arr),
		}
		return
	},
	"ecripto": func(arr []string) (e model.Event, _error bool) {
		if len(arr) < 3 {
			_error = true
			return
		}

		e = model.Event{
			Time:     deep.NewTime(),
			Key:      arr[0],
			KeyWord:  arr[1],
			Password: arr[2],
			Flags:    filterIsFlag(arr),
		}
		return
	},
	"place": func(arr []string) (e model.Event, _error bool) {
		payload := getPaylaod(arr)

		if len(arr) < 3 || payload == "" {
			_error = true
			return
		}

		e = model.Event{
			Time:    deep.NewTime(),
			Key:     arr[0],
			KeyWord: arr[1],
			Payload: payload,
		}
		return
	},
	"run":   SHORT_EVENT_KEY,
	"rm:c":  SHORT_EVENT_KEY,
	"run:m": SHORT_EVENT_KEY,
}

func SHORT_EVENT(arr []string) (e model.Event, _error bool) {
	e = model.Event{
		Time: deep.NewTime(),
		Key:  arr[0],
	}
	return
}

func SHORT_EVENT_KEY(arr []string) (e model.Event, _error bool) {
	e = model.Event{
		Time:    deep.NewTime(),
		Key:     arr[0],
		KeyWord: arr[1],
	}
	return
}
