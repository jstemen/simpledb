package simple_db

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transaction", func() {

	var (
		pTrans *Transaction
		cTrans *Transaction
	)

	BeforeEach(func() {
		pTrans = NewTransaction()
		cTrans = pTrans.New()
	})
	Describe("Get", func() {
		It("Return the value that is set", func() {
			pTrans.Set("foo", "bar")
			Expect(pTrans.Get("foo")).To(Equal("bar"))
		})
		It("Should return the value from the parent transaction if it is not in current transaction", func() {
			pTrans.Set("foo", "bar")
			Expect(cTrans.Get("foo")).To(Equal("bar"))
		})

	})

})
