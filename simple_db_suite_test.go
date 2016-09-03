package simple_db_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSimpleDb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SimpleDb Suite")
}
