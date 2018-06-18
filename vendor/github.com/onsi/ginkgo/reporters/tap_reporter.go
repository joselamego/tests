/*

TAP reporter for ginkgo

*/

package reporters

import (
	"fmt"
	"os"
	"strings"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

type TapReporter struct {
	filename string
	suite    TapTestSuite
	str      strings.Builder
}

type TapTestSuite struct {
	TestCases []TapTestCase
	Tests     int
}

type TapTestCase struct {
	Name    string
	Message string
	Details string
}

func NewTapReporter(filename string) *TapReporter {
	return &TapReporter{
		filename: filename,
	}
}

func (reporter *TapReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
	reporter.suite = TapTestSuite{
		Tests: summary.NumberOfSpecsThatWillBeRun,
	}
	fmt.Fprintf(&reporter.str, "TAP version 13\n1..%d\n", summary.NumberOfSpecsThatWillBeRun)
}

func (reporter *TapReporter) BeforeSuiteDidRun(suiteSummary *types.SetupSummary) {
}

func (reporter *TapReporter) SpecWillRun(specSummary *types.SpecSummary) {
}

func (reporter *TapReporter) SpecDidComplete(specSummary *types.SpecSummary) {
	testName := escape(strings.Join(specSummary.ComponentTexts[1:], " "))

	if specSummary.State == types.SpecStateFailed || specSummary.State == types.SpecStateTimedOut || specSummary.State == types.SpecStatePanicked {
		message := escape(specSummary.Failure.ComponentCodeLocation.String())
		details := escape(specSummary.Failure.Message)
		fmt.Fprintf(
			&reporter.str,
			"not ok %s\n\t---\n\tmessage: %s\n\tdetails: %s\n\t...\n",
			testName, message, details)
	}
	// We are handling both skipped and pending states as passed
	if specSummary.State == types.SpecStateSkipped {
		fmt.Fprintf(&reporter.str, "ok # skip %s\n", testName)
	}
	if specSummary.State == types.SpecStatePending {
		fmt.Fprintf(&reporter.str, "ok # pending %s\n", testName)
	}
	fmt.Fprintf(&reporter.str, "ok %s\n", testName)
}

func (reporter *TapReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {
}

func (reporter *TapReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	file, err := os.Create(reporter.filename)
	if err != nil {
		fmt.Printf("Failed to create Tap report file: %s\n\t%s",
			reporter.filename, err.Error())
	}
	defer file.Close()
	file.WriteString(reporter.str.String())
}
