package main

func worker(ch chan struct{}) {
	// Receive a message from the main program.
	<-ch
	println("roger")

	// Send a message to the main program.
	close(ch)
}

func main() {
	ch := make(chan struct{})
	go worker(ch)

	// Send a message to a worker.
	ch <- struct{}{}

	// Receive a message from the worker.
	<-ch
	println("roger2")
	// Output:
	// roger
	// roger
}