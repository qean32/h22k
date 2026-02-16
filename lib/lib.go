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
		key := strings.Split(
			strings.TrimSpace(command), " ")[0]
		event, _error := PARSE_EVENT(command, key)

		if !_error {
			fn := KEY_FUNCTION[key]

			if fn != nil {
				fn(event)
				deep.LOG(event)
			} else {
				fmt.Println(constants.UNDEFINED_COMMAND)
			}
		}
	}
	ENTER_COMMAND()
}

func HOF_ACCESS_ACTION(f model.EventFunction, event model.Event) {
	if ACCESS_ACTION() {
		f(event)
		return
	}

	fmt.Println(constants.STOP_COMMAND)
}

func INIT() {
	fmt.Println(constants.PROJECT_INIT)
	ENTER_COMMAND()
}

func PUSH_CYCLE() {
}

func PARSE_EVENT(command string, key string) (event model.Event, _error bool) {
	fn := KEY_PARSE[key]

	if fn == nil {
		return
	}
	event, _error = fn(strings.Split(command, " "))
	return event, _error
}

func ACCESS_ACTION() bool {
	var response string
	fmt.Print("Need access (yes/no): ")
	fmt.Scan(&response)

	if response == "yes" {

		return true
	}

	return false
}
