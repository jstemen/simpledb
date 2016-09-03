package simple_db

import (
	"io"
	"bufio"
	"fmt"
	"strings"
)

func Repl(in io.Reader, out io.Writer) {
	reader := bufio.NewReader(in)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	splits := parse(text)
	trans := NewTransaction()
	for len(splits) != 0 && splits[0] != "END" {
		cmd := splits[0]
		//fmt.Println(fmt.Sprintf("spits is %#v", splits))
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
			fmt.Fprintln(out, fmt.Sprintf("> %d",c))
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
		default:
			fmt.Fprintln(out, fmt.Sprintf("%s is not an valid command",cmd))
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
