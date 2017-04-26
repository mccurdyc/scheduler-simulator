package main

import (
	"sort"
)

// required sorting definitions
type ByCycles []Process
type ByPid []Process
type ByPriority []Process

// required sorting implementations
func (a ByCycles) Len() int             { return len(a) }
func (a ByCycles) Swap(i, j int)        { a[i], a[j] = a[j], a[i] }
func (a ByCycles) Less(i, j int) bool   { return a[i].Cycles < a[j].Cycles }
func (a ByPid) Len() int                { return len(a) }
func (a ByPid) Swap(i, j int)           { a[i], a[j] = a[j], a[i] }
func (a ByPid) Less(i, j int) bool      { return a[i].Pid < a[j].Pid }
func (a ByPriority) Len() int           { return len(a) }
func (a ByPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i, j int) bool { return a[i].Priority < a[j].Priority }

// RR algorithm; cycle through giving one cycle each
func generateFramesRR(procs []Process) [][]Process {
	var ret = [][]Process{}

	// calculate sum of cycles, processes
	var sum = 0
	var count = 0
	for _, p := range procs {
		sum += p.Cycles
		count += 1
	}

	// make enough RR passes through
	for i := 0; i <= (sum / count); i++ {
		for i := range procs {
			if procs[i].Cycles > 0 {
				procs[i].Cycles -= 1
				tmp := make([]Process, len(procs))
				copy(tmp, procs)
				ret = append(ret, tmp)
			}
		}
	}

	return ret
}

// PRR algorithm; cycle through giving one cycle each for each process in a given priority
func generateFramesPRR(procs []Process) [][]Process {
	var ret = [][]Process{}

	procs, slices := splitByPriority(procs)

	for i, _ := range slices {
		for j := 0; j < i; j++ {

			// calculate sum of cycles, processes
			var sum = 0
			var count = 0
			for _, p := range procs {
				sum += p.Cycles
				count += 1
			}

			// make enough RR passes through
			for i := 0; i <= (sum / count); i++ {
				for i := range procs {
					if procs[i].Cycles > 0 {
						procs[i].Cycles -= 1
						tmp := make([]Process, len(procs))
						copy(tmp, procs)
						ret = append(ret, tmp)
					}
				}
			}
		}
	}

	return ret
}

func splitByPriority(procs []Process) ([]Process, []int) {
	sort.Sort(ByPriority(procs))

	splitPositions := []int{}
	tmp := -1

	for i, e := range procs {
		if e.Priority != tmp {
			splitPositions = append(splitPositions, i)
		}
		tmp = e.Priority
	}
	// ret := sliceArray(procs, splitPositions)
	// return ret
	return procs, splitPositions
}

//
// func sliceArray(procs []Process, slices []int) [][]Process {
// 	ret := [][]Process{}
// 	slice := []Process{}
//
// 	for i, _ := range slices {
// 		switch i {
// 		case 0:
// 			slice = procs[:slices[i+1]]
// 		case len(slices) - 1:
// 			slice = procs[slices[i]:]
// 		default:
// 			slice = procs[slices[i]:slices[i+1]]
// 		}
// 		ret = append(ret, slice)
// 	}
// 	return ret
// }

// FCFS algorithm; just burn straight through cycles
func generateFramesFCFS(procs []Process) [][]Process {

	// define array/arrays/structs return var
	var ret = [][]Process{}

	// generate batches of frames, burning cycles
	for i := range procs {
		for procs[i].Cycles > 0 {
			procs[i].Cycles -= 1
			tmp := make([]Process, len(procs))
			copy(tmp, procs)
			ret = append(ret, tmp)
		}
	}

	return ret
}

// SJF algorithm; just sort and delegate to FCFS
func generateFramesSJF(procs []Process) [][]Process {
	sort.Sort(ByCycles(procs))
	return generateFramesFCFS(procs)
}

// SRJF algorithm; like SJF but looks at remaining cycles
func generateFramesSRJF(procs []Process) [][]Process {
	var ret = [][]Process{}

	// no real point in working on the SRJF algorithm until we sort out
	// how to get them all to consider an incoming stream of processes
	// instead of just assuming all is known beforehand

	// also would need to add fields to the proc struct

	return ret
}

// Priority algorithm; just sort and delegate to FCFS
func generateFramesP(procs []Process) [][]Process {
	sort.Sort(ByPriority(procs))
	return generateFramesFCFS(procs)
}
