package main

import (
	"os"
	"bufio"
	"strings"
	"github.com/jstemen/simple_db"
	"fmt"
)

func main() {
	in := os.Stdin
	out := os.Stdout
	reader := bufio.NewReader(in)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	splits := parse(text)
	cmd := splits[0]
	trans := simple_db.NewTransaction()
	for len(splits) != 0 && splits[0] != "END" {
		fmt.Fprintln(out, text)
		switch cmd {
		case "SET":
			trans.Set(splits[1], splits[2])
		case "GET":
			trans.Get(splits[1])
		case "UNSET":
			trans.Unset(splits[1])
		case "NUMEQUALTO":
			trans.NumEqualTo(splits[1])
		case "COMMIT":
			t, ok := trans.Commit()
			if ok {
				trans = t
			}
		case "ROLLBACK":
			t, ok := trans.Rollback()
			if ok {
				trans = t
			}
		}

	}

}

func parse(text string) (splits []string) {
	t := strings.TrimSpace(text)
	splits = strings.Split(t, " ")
	return
}
