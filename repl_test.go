package simple_db

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
	"bytes"
)

var _ = Describe("repl", func() {
	It("case one", func() {
		outb := new(bytes.Buffer)
		in := strings.NewReader("SET ex 10\nGET ex\nUNSET ex\nGET ex\nEND\n")
		Repl(in, outb)
		out := outb.String()
		Expect(out).To(Equal("SET ex 10\nGET ex\n> 10\nUNSET ex\nGET ex\n> NULL\nEND\n"))
	})

	It("case two", func() {
		outb := new(bytes.Buffer)
		in := strings.NewReader("SET a 10\nSET b 10\nNUMEQUALTO 10\nNUMEQUALTO 20\nGET ex\nSET b 30\nNUMEQUALTO 10\nEND\n")
		Repl(in, outb)
		out := outb.String()
		Expect(out).To(Equal("SET a 10\nSET b 10\nNUMEQUALTO 10\n> 2\nNUMEQUALTO 20\n> 0\nGET ex\n> NULL\nSET b 30\nNUMEQUALTO 10\n> 1\nEND\n"))
	})
})
