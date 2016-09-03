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
	Repl(os.Stdin, os.Stdout)
}
