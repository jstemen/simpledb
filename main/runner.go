package main

import (
	"bufio"
	"fmt"
	"github.com/jstemen/simple_db"
	"io"
	"os"
	"strings"
)

func main() {
	err := simple_db.Repl(os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}
}
