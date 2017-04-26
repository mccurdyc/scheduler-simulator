package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

// type Process struct {
// 	Pid      int
// 	Priority int
// 	Cycles   int
// }

func readCsv() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	var dataPath = ""
	configErr := viper.ReadInConfig()

	if configErr != nil {
		panic(configErr)
	} else {
		dataPath = viper.GetString("settings.genDataPath")
	}

	f, _ := os.Open(dataPath)
	r := csv.NewReader(bufio.NewReader(f))

	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	var processes []Process
	for _, p := range rest(records) {
		var process Process
		process.Pid, err = strconv.Atoi(p[0])
		process.Priority, err = strconv.Atoi(p[1])
		process.Cycles, err = strconv.Atoi(p[2])
		if err != nil {
			panic(err)
		}
		processes = append(processes, process)
	}
	for _, e := range processes {
		fmt.Printf("PID: %d\n", e.Pid)
		fmt.Printf("Priority: %d\n", e.Priority)
		fmt.Printf("Cycles: %d\n\n", e.Cycles)
	}
}

func rest(s [][]string) [][]string {
	return s[:][1:]
}
