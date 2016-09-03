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
			Expect(*pTrans.Get("foo")).To(Equal("bar"))
		})
		It("Should return the value from the parent transaction if it is not in current transaction", func() {
			pTrans.Set("foo", "bar")
			Expect(*cTrans.Get("foo")).To(Equal("bar"))
		})

		It("Should return nil if child transaction has been unset", func() {
			pTrans.Set("foo", "bar")
			cTrans.Unset("foo")
			var nilS *string
			Expect(cTrans.Get("foo")).To(Equal(nilS))
		})
	})
	Describe("NumEqualTo", func() {
		It("should count the correct number in a parentless transaction", func() {
			v := "waldo"
			pTrans.Set("where's", v)
			pTrans.Set("here", v)
			Expect(pTrans.NumEqualTo(v)).To(Equal(2))
		})
		It("should sum the number of occurances in parent while ignoring duplicates", func() {
			v := "waldo"
			pTrans.Set("where's", v)
			pTrans.Set("here", v)
			cTrans.Set("here", v)
			Expect(cTrans.NumEqualTo(v)).To(Equal(2))
		})

		It("should sum the number of occurances in parent", func() {
			v := "waldo"
			pTrans.Set("where's", v)
			pTrans.Set("here", v)
			cTrans.Set("there", v)
			Expect(cTrans.NumEqualTo(v)).To(Equal(3))
		})

	})

})
