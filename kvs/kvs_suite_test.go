package kvs_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKvs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kvs Suite")
}
