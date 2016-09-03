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
			Expect(cTrans.Get("foo")).To(BeNil())
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

	Describe("Commit", func() {
		It("should write child values over parent values", func() {
			pTrans.Set("foo", "lake")
			cTrans.Set("foo", "mountain")
			cTrans.Commit()
			Expect(*pTrans.Get("foo")).To(Equal("mountain"))
		})

		It("should return true when the transaction is nested", func() {
			_, ok := cTrans.Commit()
			Expect(ok).To(Equal(true))
		})

		It("should return false when the transaction is not nested", func() {
			_, ok := pTrans.Commit()
			Expect(ok).To(Equal(false))
		})

		It("should return new transaction when transaction is nested", func() {
			trans, _ := cTrans.Commit()
			Expect(trans).To(Equal(pTrans))
		})

	})

	Describe("Rollback", func() {
		It("should return parent when there is one", func() {
			t, ok := cTrans.Rollback()
			Expect(t).To(Equal(pTrans))
			Expect(ok).To(Equal(true))
		})
		It("should return nil when there is no parent", func() {
			t, ok := pTrans.Rollback()
			Expect(t).To(BeNil())
			Expect(ok).To(Equal(false))
		})

		It("should disconnect the parent from the child", func() {
			cTrans.Rollback()
			Expect(pTrans.Child).To(BeNil())
			Expect(cTrans.Parent).To(BeNil())
		})

	})

})
