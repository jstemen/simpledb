package main

import (
	"os"
	"bufio"
	"strings"
	"github.com/jstemen/simple_db"
	"fmt"
	"io"
)

func main() {
	err := simple_db.Repl(os.Stdin, os.Stdout)
	if err != nil{
		panic(err)
	}
}
