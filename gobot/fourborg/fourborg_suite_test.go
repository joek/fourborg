package fourborg_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFourBorg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FourBorg Suite")
}
