package deep

import (
	"bufio"
	"fmt"
	"main/constants"
	"main/model"
	"os"
	"strings"
	"time"
)

func LOG(event model.Event) {
	PushToFile(constants.LOG_PATH, fmt.Sprintf("%#v", event))
}

func ITERATION_CYCLE() {
}

var CALLSTACK = []string{}

func ReadFile(path string) []string {
	file, err := os.Open(constants.Root + path)
	scanner := bufio.NewScanner(file)

	if err != nil {
		return nil
	}
	defer file.Close()

	var data []string
	for scanner.Scan() {
		data =
			append(data, scanner.Text())
	}
	return data
}

func WriteFile(data string, path string) bool {
	file, err := os.Create(constants.Root + path)

	if err != nil {
		return false
	}
	defer file.Close()

	file.WriteString(data)
	return true
}

func PushToFile(path string, newText string) bool {
	return WriteFile(strings.Join(
		append(ReadFile(path), newText+"\n"), " \n"), path)
}

func CreateFile(path string) {
	file, _ := os.Create(constants.Root + path)
	defer file.Close()
}

func GenerateMaster() {

}

func ClearFile(path string) {
	WriteFile("", path)
}

func NewTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
