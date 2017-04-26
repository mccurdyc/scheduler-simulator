package main

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Process struct {
	Pid      int
	Priority int
	Cycles   int
}

func generateProcesses(n int, file string) []Process {
	rand.Seed(time.Now().UTC().UnixNano()) // otherwise always same random values

	const maxPriority = 8
	const maxCycles = 10
	processes := make([]Process, n)

	i := 0
	for i < n {
		processes[i].Pid = i
		processes[i].Priority = rand.Intn(maxPriority)
		processes[i].Cycles = rand.Intn(maxCycles)
		i++
	}
	writeToFile(file, processes)
	return processes
}

func writeToFile(file string, processes []Process) {
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Write([]string{"PID", "Priority", "Cycles"})
	for _, p := range processes {
		strPid := strconv.Itoa(p.Pid)
		strPriority := strconv.Itoa(p.Priority)
		strCycles := strconv.Itoa(p.Cycles)

		var arr = []string{strPid, strPriority, strCycles}
		err = w.Write(arr)
		if err != nil {
			panic(err)
		}
	}
	w.Flush()
}
