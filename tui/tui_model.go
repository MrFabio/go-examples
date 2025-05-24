package tui

import (
	datatype "go-examples/datatypes"
	functions "go-examples/functions"
	lang "go-examples/lang"
	concurrency "go-examples/lang/concurrency"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
)

func InitialModel() model {
	items := []list.Item{
		item{title: "Rune Examples", run: datatype.RuneExamples, file: "datatypes/runes.go", funcName: "RuneExamples"},
		item{title: "Strings Examples", run: datatype.StringsExamples, file: "datatypes/strings.go", funcName: "StringsExamples"},
		item{title: "Ints Examples", run: datatype.IntsExamples, file: "datatypes/ints.go", funcName: "IntsExamples"},
		item{title: "Time Examples", run: functions.TimeExamples, file: "functions/time.go", funcName: "TimeExamples"},
		item{title: "Loops Examples", run: functions.LoopsExamples, file: "functions/loops.go", funcName: "LoopsExamples"},
		item{title: "Arrays Examples", run: datatype.ArraysExamples, file: "datatypes/arrays.go", funcName: "ArraysExamples"},
		item{title: "Slices Examples", run: datatype.SlicesExamples, file: "datatypes/slices.go", funcName: "SlicesExamples"},
		item{title: "Calls Examples", run: functions.CallsExamples, file: "functions/calls.go", funcName: "CallsExamples"},
		item{title: "Functions Examples", run: functions.FunctionsExamples, file: "functions/functions.go", funcName: "FunctionsExamples"},
		item{title: "IO Examples", run: functions.IoExamples, file: "functions/io.go", funcName: "IoExamples"},
		item{title: "CLI Examples", run: functions.CliExamples, file: "functions/cli.go", funcName: "CliExamples"},
		item{title: "Maps Examples", run: datatype.MapsExamples, file: "datatypes/maps.go", funcName: "MapsExamples"},
		item{title: "Generics Examples", run: datatype.GenericsExamples, file: "datatypes/generics.go", funcName: "GenericsExamples"},
		item{title: "Structs Examples", run: datatype.StructsExamples, file: "datatypes/structs.go", funcName: "StructsExamples"},
		item{title: "Composition Examples", run: datatype.CompositionExamples, file: "datatypes/composition.go", funcName: "CompositionExamples"},
		item{title: "Accessors Examples", run: lang.AccessorsExamples, file: "lang/accessors.go", funcName: "AccessorsExamples"},
		item{title: "Interfaces Examples", run: lang.InterfacesExamples, file: "lang/interfaces.go", funcName: "InterfacesExamples"},
		item{title: "Routines Examples", run: concurrency.RoutinesExamples, file: "lang/concurrency/routines.go", funcName: "RoutinesExamples"},
		item{title: "Simple Channels Examples", run: concurrency.SimpleChannelsExamples, file: "lang/concurrency/channels.go", funcName: "SimpleChannelsExamples"},
		item{title: "Channels with Mutex Examples", run: concurrency.ChannelsWithMutexExamples, file: "lang/concurrency/channels.go", funcName: "ChannelsWithMutexExamples"},
		item{title: "Closures Examples", run: lang.ClosuresExamples, file: "lang/closures.go", funcName: "ClosuresExamples"},
		item{title: "Recursion Examples", run: lang.RecursionExamples, file: "lang/recursion.go", funcName: "RecursionExamples"},
		item{title: "Regex Examples", run: lang.RegexExamples, file: "lang/regex.go", funcName: "RegexExamples"},
	}

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	l := list.New(items, delegate, 0, 0)
	l.Title = "Go Examples"
	l.SetShowHelp(true)

	outputVP := viewport.New(0, 0)
	codeVP := viewport.New(0, 0)

	return model{list: l, mode: modeList, outputVP: outputVP, codeVP: codeVP}
}
