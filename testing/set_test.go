package testing_test

import (
	// using "." here so that we directly use Describe and other ginko fucntions

	// upgrade to ginkgo v2 https://onsi.github.io/ginkgo/MIGRATING_TO_V2

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// READ on https://github.com/onsi/ginkgo

/*
func TestTesting2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testing Suite2")
} */

var _ = Describe("Set", func() { //Describe
	Describe("Emptiness", func() { //Describe
		/* AfterEach(func() {
			log.Println("AfterTest: p is empty")
		}) */

		AfterEach(func() {
			if CurrentSpecReport().Failed() {
				GinkgoWriter.Println("printing to stdout failed report")
			}
			if !CurrentSpecReport().Failed() {
				GinkgoWriter.Println("printing to stdout pass report")
			}
		})
		Context("When the set does not contain items", func() { //Context
			It("Should be empty", func() { //It
				// Use of By
				By("if set is empty", func() {
					set := ""
					Expect(set).To(BeEquivalentTo(""))
				})
			})

		})
	})
})

/*Other Examples use of ERROR and EVENTUALLY.
 When("the library does not have the book in question", func() {
        It("tells the reader the book is unavailable", func() {
            err := valjean.Checkout(library, "Les Miserables")
            Expect(error).To(MatchError("Les Miserables is not in the library catalog"))
        })
		// use of Eventually
		It("can connect to the server", func() {
      	Eventually(client.Connect).Should(Succeed())
    	})
*/
