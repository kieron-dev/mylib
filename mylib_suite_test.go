package mylib_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMylib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mylib Suite")
}
