package simple_db_test

import (
	. "github.com/jstemen/simple_db"
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("BidirectionalMap", func() {

	Describe("Set", func() {
		var (
			m *BidirectionalMap
		)
		BeforeEach(func() {
			m = NewBidirectionalMap()
		})
		It("should set the value", func() {
			m.Set("foo", "bar")
			val, _ := m.Get("foo")
			Expect(val).To(Equal("bar"))
		})
		It("should set the reverse value", func() {
			m.Set("foo", "bar")
			vals, _ := m.GetKeysFromValue("bar")
			Expect(vals["foo"]).To(Equal(true))
		})
		It("should delete old values for key from inverse map", func() {
			m.Set("foo", "bar")
			m.Set("foo", "pumpkin")
			vals, ok := m.GetKeysFromValue("bar")

			Expect(len(vals)).To(Equal(0))
			Expect(ok).To(Equal(false))
		})
	})
})
