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
		inputs := strings.Fields(input)

		if len(inputs) == 0 {
			continue
		}

		var (
			cmd  string
			arg1 string
			arg2 string

			err error
		)

		for i, in := range inputs {
			switch i {
			case 0:
				cmd = in
			case 1:
				arg1 = in
			case 2:
				arg2 = in
			case 3:
				break
			}
		}

		result, err := handle(store, cmd, arg1, arg2)
		if err != nil {
			printMsg(err)
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
