package kvs

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Proxy Tests", func() {
	var store *KVS
	BeforeEach(func() {
		store = NewStore(WithMaxRAMSize(4))
	})

	AfterEach(func() {
		store = nil
	})

	When("max capacity of the in-memory store is reached", func() {
		It("should evict the least recently used item", func() {
			respCode := store.InsertItem("user1", 1)
			Expect(respCode).To(Equal(1))

			respCode = store.InsertItem("user2", 2)
			Expect(respCode).To(Equal(1))

			respCode = store.InsertItem("user3", 3)
			Expect(respCode).To(Equal(1))

			respCode = store.InsertItem("user4", 4)
			Expect(respCode).To(Equal(1))

			respCode = store.InsertItem("user5", 5)
			Expect(respCode).To(Equal(1))

			value := store.GetItem("user1")
			Expect(value).To(Equal("NULL"))
		})
	})

	When("it receives a insertion request", func() {
		It("it should insert the item in the cache successfully", func() {
			respCode := store.InsertItem("user", 123)
			Expect(respCode).To(Equal(1))
		})
	})

	When("it receives a read request", func() {
		It("it should read the item in the cache successfully", func() {
			respCode := store.InsertItem("user", 123)
			Expect(respCode).To(Equal(1))

			value := store.GetItem("user")
			Expect(value).To(Equal(123))
		})
	})

	When("it receives a update request", func() {
		It("it should update the item in the cache successfully", func() {
			respCode := store.InsertItem("user", 123)
			Expect(respCode).To(Equal(1))

			value := store.GetItem("user")
			Expect(value).To(Equal(123))

			respCode = store.UpdateItem("user", 456)
			Expect(respCode).To(Equal(1))

			value = store.GetItem("user")
			Expect(value).To(Equal(456))
		})
	})

	When("it receives a delete request", func() {
		It("it should delete the item in the cache successfully", func() {
			respCode := store.InsertItem("user", 123)
			Expect(respCode).To(Equal(1))

			value := store.GetItem("user")
			Expect(value).To(Equal(123))

			respCode = store.DeleteItem("user")
			Expect(respCode).To(Equal(1))

			value = store.GetItem("user")
			Expect(value).To(Equal("NULL"))
		})
	})
})
