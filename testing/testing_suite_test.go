package testing_test

import (
	"fmt"
	"log"
	"os"

	// upgrade to ginkgo v2 https://onsi.github.io/ginkgo/MIGRATING_TO_V2
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	testing1 "ginkgotutorial/testing"
	"testing"
)

func TestTesting1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testing Suite1")
}

// ReportAfterSuite is called exactly once at the end of the suite after any AfterSuite nodes have run
// IMPORTANT ReportAfterSuite only runs on process #1 and receives a Report that aggregates the SpecReports from all processes.
var _ = ReportAfterSuite("Custom report ReportAfterSuite:", func(report Report) {
	// process report
	log.Println("Printing All the reports after suite")
	for _, specReport := range report.SpecReports {
		fmt.Println("printing report after suite")
		fmt.Printf("ReportAFterSuite: %s | %s\n", specReport.FullText(), specReport.State)
	}
	//customFormat := fmt.Sprintf("Time taken %v, pre-run stats %v , suiteDesc %s, flakedstate %d", report.RunTime, report.PreRunStats, report.SuiteDescription, report.SpecReports.CountOfFlakedSpecs())
	//fmt.Println(customFormat)
	AddReportEntry("Report enteries: ReportEntry includes the passed in name as well as the time and source location at which AddReportEntry was called", func() {
		log.Println("additional information by AddReportEntry")
	})
	// We can also print the AfterSuiteReport in a file.
	f, _ := os.Create("report.custom")
	for _, specReport := range report.SpecReports {
		fmt.Fprintf(f, "%s | %s\n", specReport.FullText(), specReport.State)
	}
	f.Close()
})

var _ = Describe("1PersonIsChild()", func() {
	/* AfterSuite(func() {
		log.Println("After sutie")
	}) */
	ReportBeforeEach(func(report SpecReport) {
		customFormat := fmt.Sprintf("%s | %s", report.FullText(), report.ReportEntries)
		fmt.Println("Report Before Each", customFormat)
	})
	ReportAfterEach(func(report SpecReport) {
		customFormat := fmt.Sprintf("Custom reportAfterEach: %s | Fulltext: %s", report.State, report.FullText())
		fmt.Println(customFormat)
		AddReportEntry("Time", report.RunTime)
		fmt.Println("Report Entries under afterEachReport", report.ReportEntries)
	})
	Context("When the person is child", func() {
		It("returns True", func() {
			person := testing1.Person{
				Age: 10,
			}
			response := person.IsChild()
			Expect(response).To(BeTrue()) // note here we can also write Expect(person.IsChild()).To(BeTrue())
			// So IsChild mostly user defined method in the file being tested

		})
		fmt.Println("Current spec report example:", CurrentSpecReport())
	})
	Context("When the person is NOT a child", func() {
		BeforeEach(func() {
			log.Println("BeforeTest: Person is not a child")
		})

		AfterEach(func() {
			log.Println("AfterTest: Person after")
		})
		It("returns True", func() {
			person := testing1.Person{
				Age: 20,
			}
			response := person.IsChild()
			Expect(response).To(BeFalse())
		})
	})
	Context("When the person is empty", func() {
		It("returns True", func() {

			response := ""
			Expect(response).To(BeEmpty())
		})
	})
	Context("When the person is invalid and a number", func() {
		BeforeEach(func() {
			log.Println("BeforeTest numericalcheck: Person is invalid")
		})
		It("returns True", func() {

			response := 1.0
			Expect(response).Should(BeNumerically("==", 1))
		})
	})
	Describe("2PersonIsChild()", func() {
		Context("context", func() {
			It("return true", func() {
				var a int
				Expect(a).To(BeZero())
			})
		})
		AfterEach(func() {
			log.Println("a is zero")
		})
	})
})
