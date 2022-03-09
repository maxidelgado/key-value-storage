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
		var args [3]string
		fmt.Print("\n> ")
		rawInput, _ := reader.ReadString('\n')
		inputs := strings.Fields(rawInput)
		for i, input := range inputs {
			if i == 2 {
				break
			}
			args[i] = input
		}

		result, err := handle(store, args[0], args[1], args[2])
		if err != nil {
			printMsg("ERROR: " + err.Error())
		} else if result != nil {
			printMsg(result)
		}
	}
}

func handle(store kvs.KVS, cmd, arg1, arg2 string) (result interface{}, err error) {
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
