package lib

import (
	"bufio"
	"fmt"
	"main/constants"
	"main/deep"
	"main/model"
	"os"
	"strings"
)

var READER = bufio.NewReader(os.Stdin)

func ENTER_COMMAND() {
	print("> ")
	command, _ := READER.ReadString('\n')

	if len(command) > 1 {
		trimString := strings.TrimSpace(command)
		key := strings.Split(trimString, " ")[0]
		e, _error := PARSE_EVENT(trimString, key)

		if !_error {
			fn := KEY_FUNCTION[key]

			if fn != nil {
				fn(e)
				deep.LOG(e)
			} else {
				fmt.Println(constants.UNDEFINED_COMMAND)
			}
		}
	}
	ENTER_COMMAND()
}

func INIT() {
	constants.INIT_ROOT()
	fmt.Println(constants.Root)
	fmt.Println(constants.PROJECT_INIT)
	deep.SET_DATA()
	ENTER_COMMAND()
}

func PUSH_CYCLE() {
}

func PARSE_EVENT(command string, key string) (e model.Event, _error bool) {
	fn := KEY_PARSE[key]

	if fn == nil {
		e, _error = SHORT_EVENT(strings.Split(command, " "))
		return
	}
	e, _error = fn(strings.Split(command, " "))
	return e, _error
}
