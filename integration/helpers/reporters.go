package helpers

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/ginkgo/reporters/stenographer"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

func GetPRBuilderReporter() ginkgo.Reporter {
	outputDir := os.Getenv("PR_BUILDER_OUTPUT_DIR")

	// TODO: consider returning an error
	if outputDir == "" {
		return nil
	}

	prBuilderReporter := NewPRBuilderReporter(outputDir)
	return prBuilderReporter
}

type PRBuilderReporter struct {
	nestedReporter *reporters.DefaultReporter
	debugFile      *os.File
	stenographer   stenographer.Stenographer
	specSummaries  []*types.SpecSummary
}

func NewPRBuilderReporter(outputDir string) *PRBuilderReporter {

	outputFile := filepath.Join(outputDir, strconv.Itoa(ginkgo.GinkgoParallelNode()))

	f, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	reporter := &PRBuilderReporter{debugFile: f, stenographer: stenographer.New(false, true, f)}
	c := config.DefaultReporterConfig
	c.Succinct = true
	c.Verbose = false
	c.FullTrace = false
	reporter.nestedReporter = reporters.NewDefaultReporter(c, reporter.stenographer)

	return reporter
}

func (reporter *PRBuilderReporter) SpecSuiteWillBegin(conf config.GinkgoConfigType, summary *types.SuiteSummary) {
	//reporter.nestedReporter.SpecSuiteWillBegin(conf, summary)
	//reporter.debugFile.Sync()
}

func (reporter *PRBuilderReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {
	reporter.nestedReporter.BeforeSuiteDidRun(setupSummary)
	reporter.debugFile.Sync()
}

func (reporter *PRBuilderReporter) SpecWillRun(specSummary *types.SpecSummary) {
	reporter.nestedReporter.SpecWillRun(specSummary)
	reporter.debugFile.Sync()
}

func (reporter *PRBuilderReporter) SpecDidComplete(specSummary *types.SpecSummary) {
	//reporter.nestedReporter.SpecDidComplete(specSummary)
	reporter.specSummaries = append(reporter.specSummaries, specSummary)
	reporter.debugFile.Sync()
}

func (reporter *PRBuilderReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {
	reporter.nestedReporter.AfterSuiteDidRun(setupSummary)
	reporter.debugFile.Sync()
}

func (reporter *PRBuilderReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	reporter.stenographer.SummarizeFailures(reporter.specSummaries)
	reporter.debugFile.Sync()
}
