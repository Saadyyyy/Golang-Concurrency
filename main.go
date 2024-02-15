package main

import (
	"fmt"
	"time"
)

// Ping-Pong sederhana
// buat sebuah aplikasi sederhana dengan kriteria
// 2 pemain ping-pong saling membalikan bola
// pemainan berhenti setelah 1 detik
// gunakan konkurensi dengan 1 channel
func main() {
	table := make(chan *Ball)

	go Player("laode", table)
	go Player("Saady", table)

	table <- new(Ball)
	time.Sleep(1 * time.Second)
	<-table
}

type Ball struct {
	hits int
}

func Player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, "Hits The Ball", ball.hits)
		time.Sleep(50 * time.Millisecond)
		table <- ball
	}
}
