package mylib_test

import (
	"github.com/kieron-pivotal/mylib"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mylib", func() {
	It("returns something", func() {
		Expect(mylib.HaveIt()).ToNot(BeEmpty())
	})
})
