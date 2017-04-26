package main

import (
	"github.com/gizak/termui"
)

func setupInterface(procs []Process) {
	defer termui.Close()
	defer termui.Loop()

	// initialize termui
	err := termui.Init()
	if err != nil {
		panic(err)
	}

	// keybindings pane
	s := " q => Exits Simulator\n"
	s += " r => Launch RR Simulation\n"
	s += " R => Launch PRR Simulation\n"
	s += " s => Launch SJF Simulation\n"
	s += " S => Launch SRJF Simulation\n"
	s += " f => Launch FCFS Simulation\n"
	s += " p => Launch Priority Simulation\n"
	p := termui.NewPar(s)
	p.Height = 8
	p.Width = 45
	p.TextFgColor = termui.ColorWhite
	p.BorderLabel = "Keybindings"
	p.BorderFg = termui.ColorCyan
	termui.Render(p)
	//
	// // calculate layout
	// termui.Body.AddRows(
	// 	termui.NewRow(
	// 		termui.NewCol(6, 0, p),
	// 		termui.NewCol(6, 0, stats)),
	// 	termui.NewRow(
	// 		termui.NewCol(3, 0, bc)))
	//
	// termui.Body.Align()
	// termui.Render(termui.Body)

	// exit simulator and termui
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	// kickoff shortest job first algorithm simulation
	termui.Handle("/sys/kbd/s", func(termui.Event) {
		stepThrough(generateFramesSJF(procs))
	})
	// kickoff shortest remaining job first algorithm simulation
	termui.Handle("/sys/kbd/S", func(termui.Event) {
		// stepThrough(generateFramesSRJF(procs))
	})
	// kickoff round robin algorithm simulation
	termui.Handle("/sys/kbd/r", func(termui.Event) {
		stepThrough(generateFramesRR(procs))
	})
	// kickoff priority round robin algorithm simulation
	termui.Handle("/sys/kbd/R", func(termui.Event) {
		stepThrough(generateFramesPRR(procs))
	})
	// kickoff first come first serve algorithm simulation
	termui.Handle("/sys/kbd/f", func(termui.Event) {
		stepThrough(generateFramesFCFS(procs))
	})
	// kickoff priority based algorithm simulation
	termui.Handle("/sys/kbd/p", func(termui.Event) {
		stepThrough(generateFramesP(procs))
	})
}
