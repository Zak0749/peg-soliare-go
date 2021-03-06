package main

import (
	"encoding/json"
	"fmt"
	moves "main/get-moves"
	"main/starting-board"
	"main/unique"
	"os"
	"time"
)

func main() {
	data := [][][45]bool{}
	main := make(chan [][45]bool, 44)
	now := time.Now()

	go computeLayer([][45]bool{starting.Board}, main)

	for cur := range main {
		fmt.Println(len(data)+1, "layers computed items:", len(cur), "time", time.Since(now))
		now = time.Now()
		data = append(data, cur)
	}
	close(main)

	file, _ := json.MarshalIndent(data, "", " ")

	os.WriteFile("data.json", file, 0644)
}

func computeLayer(boards [][45]bool, main chan<- [][45]bool) {
	layer := [][45]bool{}
	jobs := make(chan [45]bool)
	results := make(chan [][45]bool, len(boards))

	for i := 0; i < 500; i++ {
		go worker(jobs, results)
	}

	for _, board := range boards {
		jobs <- board
	}

	close(jobs)

	for j := 0; j < len(boards); j++ {
		res := <-results
		layer = append(layer, res...)
	}
	close(results)

	if len(layer) == 0 {
		return
	}

	main <- layer

	unique := unique.Boards(layer)

	go computeLayer(unique, main)
}

func worker(jobs <-chan [45]bool, results chan<- [][45]bool) {
	for n := range jobs {
		results <- moves.Get(n)
	}
}
