package deep

import (
	"bufio"
	"fmt"
	"main/constants"
	"main/model"
	"os"
	"os/exec"
	"slices"
	"strings"
	"time"
)

func LOG(e model.Event) {
	if slices.IndexFunc(e.Flags, func(item string) bool {
		return strings.TrimSpace(item) == "-nl"
	}) == -1 {
		PushToFile(constants.LOG_PATH, fmt.Sprintf("%#v", e))
	}
}

func ITERATION_CYCLE() {
}

var CALLSTACK = []string{}
var TMP_DATA = [][]string{}
var TMP_COMMANDS = [][]string{}

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

func SET_DATA() {
	if len(TMP_COMMANDS) == 0 {
		strs := (ReadFile(constants.COMMAND_PATH))
		var tmpArr [][]string

		for i := 0; i < len(strs); i++ {
			tmp := strings.Split(strs[i], " ")
			tmpArr = append(tmpArr, []string{tmp[0], strings.Join(tmp[1:], " ")})
		}
		TMP_COMMANDS = tmpArr
	}
	if len(TMP_DATA) == 0 {
		strs := (ReadFile(constants.DATA_PATH))
		var tmpArr [][]string

		for i := 0; i < len(strs); i++ {
			tmp := strings.Split(strs[i], " ")
			tmpArr = append(tmpArr, []string{tmp[0], strings.Join(tmp[1:], " ")})
		}
		TMP_DATA = tmpArr
	}
}

func MatrixToArrayString(matrix [][]string) []string {
	var tmpArr []string

	for i := 0; i < len(matrix); i++ {
		tmpArr = append(tmpArr, strings.Join(matrix[i], " "))
	}

	return tmpArr
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

func DECORATOR_ACCESS_ACTION(f model.EventFunction) model.EventFunction {
	return func(e model.Event) {

		if ACCESS_ACTION() {
			f(e)
		} else {
			fmt.Println(constants.STOP_COMMAND)
		}
	}
}

func RunCommand(command string) {
	cmd := exec.Command("CMD.exe", "/C", command)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при запуске команды:", err)
		return
	}
}
