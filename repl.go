package simple_db

import (
	"io"
	"bufio"
	"fmt"
	"strings"
)

func Repl(in io.Reader, out io.Writer) (err error) {
	reader := bufio.NewReader(in)
	var text string
	text, err = reader.ReadString('\n')
	if err != nil {
		return
	}
	splits := parse(text)
	trans := NewTransaction()
	for len(splits) != 0 && splits[0] != "END" {
		cmd := splits[0]
		fmt.Fprintf(out, text)
		switch cmd {
		case "SET":
			trans.Set(splits[1], splits[2])
		case "GET":
			s := trans.Get(splits[1])

			fmt.Fprintf(out, "> ")
			if s == nil {
				fmt.Fprintln(out, "NULL")
			}else {
				fmt.Fprintln(out, *s)
			}
		case "UNSET":
			trans.Unset(splits[1])
		case "NUMEQUALTO":
			c := trans.NumEqualTo(splits[1])
			fmt.Fprintln(out, fmt.Sprintf("> %d", c))
		case "COMMIT":
			t, ok := trans.Commit()
			if ok {
				trans = t
			}else {
				fmt.Fprintln(out, "> NO TRANSACTION")
			}
		case "ROLLBACK":
			t, ok := trans.Rollback()
			if ok {
				trans = t
			}else {
				fmt.Fprintln(out, "> NO TRANSACTION")
			}
		case "BEGIN":
			trans = trans.New()
		default:
			fmt.Fprintln(out, fmt.Sprintf("%s is not an valid command", cmd))
		}
		text, err = reader.ReadString('\n')
		splits = parse(text)
	}
	fmt.Fprint(out, text)
}

func parse(text string) (splits []string) {
	t := strings.TrimSpace(text)
	splits = strings.Split(t, " ")
	return
}
