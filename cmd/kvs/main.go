package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"key-value-storage/internal/app/kvs"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	store := kvs.New()
	for {
		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')
		result, err := handle(store, input)
		if err != nil {
			printMsg("ERROR: " + err.Error())
		} else if result != nil {
			printMsg(result)
		}
	}
}

func parseInput(rawInput string) (string, string, string) {
	var args [3]string
	inputs := strings.Fields(rawInput)
	for i, input := range inputs {
		if i == 3 {
			break
		}
		args[i] = input
	}
	return args[0], args[1], args[2]
}

func handle(store kvs.KVS, input string) (result interface{}, err error) {
	cmd, arg1, arg2 := parseInput(input)
	switch strings.ToUpper(cmd) {
	case "SET":
		store.Set(arg1, arg2)
	case "GET":
		result, err = store.Get(arg1)
	case "DELETE":
		err = store.Delete(arg1)
	case "COUNT":
		result = store.Count(arg1)
	case "BEGIN":
		store.Begin()
	case "ROLLBACK":
		err = store.Rollback()
	case "COMMIT":
		err = store.Commit()
	default:
		err = errors.New("unrecognised command " + cmd)
	}
	return
}

func printMsg(val interface{}) {
	fmt.Printf("/> %v\n", val)
}
