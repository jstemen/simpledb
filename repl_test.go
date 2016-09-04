package simple_db

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
	"bytes"
)

var _ = Describe("repl", func() {
	Context("Sample example cases: ", func() {
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

		It("case three", func() {
			outb := new(bytes.Buffer)
			in := strings.NewReader("BEGIN\nSET a 30\nBEGIN\nSET a 40\nCOMMIT\nGET a\nROLLBACK\nCOMMIT\nEND\n")
			Repl(in, outb)
			out := outb.String()
			Expect(out).To(Equal("BEGIN\nSET a 30\nBEGIN\nSET a 40\nCOMMIT\nGET a\n> 40\nROLLBACK\n> NO TRANSACTION\nCOMMIT\n> NO TRANSACTION\nEND\n"))
		})

		It("case four", func() {
			outb := new(bytes.Buffer)
			in := strings.NewReader("BEGIN\nSET a 30\nBEGIN\nSET a 40\nCOMMIT\nGET a\nROLLBACK\nCOMMIT\nEND\n")
			Repl(in, outb)
			out := outb.String()
			Expect(out).To(Equal("BEGIN\nSET a 30\nBEGIN\nSET a 40\nCOMMIT\nGET a\n> 40\nROLLBACK\n> NO TRANSACTION\nCOMMIT\n> NO TRANSACTION\nEND\n"))
		})


		It("case five", func() {
			outb := new(bytes.Buffer)
			in := strings.NewReader("SET a 50\nBEGIN\nGET a\nSET a 60\nBEGIN\nUNSET a\nGET a\nROLLBACK\nGET a\nCOMMIT\nGET a\nEND\n")
			Repl(in, outb)
			out := outb.String()
			Expect(out).To(Equal("SET a 50\nBEGIN\nGET a\n> 50\nSET a 60\nBEGIN\nUNSET a\nGET a\n> NULL\nROLLBACK\nGET a\n> 60\nCOMMIT\nGET a\n> 60\nEND\n"))
		})

		It("case six", func() {
			outb := new(bytes.Buffer)
			in := strings.NewReader("SET a 10\nBEGIN\nNUMEQUALTO 10\nBEGIN\nUNSET a\nNUMEQUALTO 10\nROLLBACK\nNUMEQUALTO 10\nCOMMIT\nEND\n")
			Repl(in, outb)
			out := outb.String()
			Expect(out).To(Equal("SET a 10\nBEGIN\nNUMEQUALTO 10\n> 1\nBEGIN\nUNSET a\nNUMEQUALTO 10\n> 0\nROLLBACK\nNUMEQUALTO 10\n> 1\nCOMMIT\nEND\n"))
		})

	})
})
