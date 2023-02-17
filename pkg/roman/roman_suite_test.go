package roman_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRoman(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Roman Suite")
}
