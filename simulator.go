package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/gizak/termui"
)

func main() {

	// generate process data
	procs := generateProcesses(20, "../data/sampledata.csv")

	// setup the interface
	setupInterface(procs)
}

func stepThrough(frames [][]Process) {

	// define variables outside loop
	stat1, stat2 := 0, 0

	// given frames, step through them
	for i := range frames {

		// calculating statistics
		stat1 += 3
		stat2 += 4

		// waiting for some time
		time.Sleep(1000 * time.Millisecond)

		// then pushing the data to refreshInterface
		refreshInterface(frames[i], stat1, stat2)
	}
}

func refreshInterface(procs []Process, stat1 int, stat2 int) {

	bc := termui.NewBarChart()
	bc.Y = 10

	data := []int{}
	sort.Sort(ByPid(procs))
	for i := range procs {
		data = append(data, procs[i].Cycles)
	}

	labels := []string{
		"P0", "P1", "P2", "P3", "P4",
		"P5", "P6", "P7", "P8", "P9",
		"P10", "P11", "P12", "P13", "P14",
		"P15", "P16", "P17", "P18", "P19"}

	bc.BorderLabel = "Processes"
	bc.Data = data
	bc.Width = 90
	bc.Height = 15
	bc.SetMax(20)
	bc.DataLabels = labels
	bc.TextColor = termui.ColorGreen
	bc.BarColor = termui.ColorRed
	bc.NumColor = termui.ColorYellow

	// render
	termui.Render(bc)

	// statistics pane
	s := fmt.Sprintf("Stat1: %d\nStat2: %d\n", stat1, stat2)
	st := termui.NewPar(s)
	st.X = 45
	st.Height = 8
	st.Width = 45
	st.TextFgColor = termui.ColorWhite
	st.BorderLabel = "Calculated Statistics"
	st.BorderFg = termui.ColorCyan
	termui.Render(st)
}
