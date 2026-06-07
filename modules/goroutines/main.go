package main

// "fmt"
// "time"

func main() {
	// goroutines are used mainly in multithreading. We use the keyword go to acheive this
	// Eg: In this loop everything runs normally and the Iterations are printed sequentially in the console
	// fmt.Println("Normal loop execution:")
	// for i := range 10 {
	// 	fmt.Println("Iteration:", i)
	// }

	// For this function as the print statements are ran in multiple threads. The Iteration sequence is random and based on the threads completion rate
	// This output is also not printed in the console without sleeping because the main function also runs on a thread. The main function finishes first before every loop is completed so there is no output for this
	// fmt.Println("Multithreaded loop execution (The result is different and not sequentially achieved):")
	// for idx := range 10 {
	// 	go func(i int) {
	// 		fmt.Println("Iteration:", i)
	// 	}(idx)
	// }

	// use sleep for a second to see the multithreading result
	// time.Sleep(time.Second)
	// When setting sleep the function runs successfully
	// Having a second loop also does not help without waiting for the multithreading to work because before the background threads get fired off the main function exits already
	// In applications the programmer does not have any idea of how long should the program wait so waitgroups are implemented
	// Function for waitgroup demo
	// waitGrp()
	UnbufferedChannelDemo()
}
