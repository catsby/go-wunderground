package main

import (
	"fmt"
	wu "github.com/logank/go-wunderground"
	wt "github.com/logank/go-wunderground/wutest"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestTemplatesUS(t *testing.T) {
	testTemplates(t, UCustomary)
}

func TestTemplatesSI(t *testing.T) {
	testTemplates(t, UMetric)
}

func testTemplates(t *testing.T, u UnitSelect) {
	ar := wt.GetTestApiResponse()
	pr := wt.GetTestPlanner()

	te, err := TemplateEngine(u)
	if err != nil {
		t.Errorf("failed to parse alerts: %s", err)
	}

	run := func(name string, v interface{}) {
		var w io.Writer = ioutil.Discard
		if testing.Verbose() {
			w = os.Stderr
		}
		fmt.Fprintln(w, name, ">>>")
		if err := te.Lookup(name).Execute(w, v); err != nil {
			t.Errorf("failed template %s: %s", name, err)
		}
		fmt.Fprintln(w, "<<<")
	}

	run("Alerts", ar.Alerts)
	run("Alerts", []wu.Alerts{})
	run("Almanac", ar.Almanac)
	run("CurrentObservation", ar.CurrentObservation)
	run("Forecast", ar.Forecast)
	run("History", ar.History)
	run("Location", ar.Location)
	run("MoonPhase", ar.MoonPhase)
	run("Tide", ar.Tide)
	noTide := *ar.Tide
	noTide.TideSummary = noTide.TideSummary[:0]
	run("Tide", noTide)
	run("Trip", pr.Trip)
	run("ApiResponse", ar)
	run("ApiResponse", pr)
}
