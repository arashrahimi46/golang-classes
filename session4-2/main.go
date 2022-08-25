package main

import (
	"fmt"
	"sync"
)
type Person struct {
	Name string
	Age int
}

func main()  {
	namesChannel := make(chan string)
	personChannel := make(chan Person)
	var namesArray = []string{
		"arash", "hamed" , "ali" , "zahra" , "maryam",
	}
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go readFromChannel(namesChannel, &wg)
	go writeInChannel(namesArray , namesChannel)
	wg.Add(1)
	go readPersonChannel(personChannel , &wg)
	for i := 0 ; i < 3 ; i++ {
		personChannel <- Person{
			Name: "Arash",
			Age:  i,
		}
	}

	close(personChannel)
	wg.Wait()
}

func readPersonChannel(channel chan Person, w *sync.WaitGroup) {
	defer w.Done()
	for person := range channel {
		fmt.Println(person)
	}
}

func writeInChannel(array []string, channel chan string) {
	for _, name := range array {
		channel <- name
	}
	close(channel)
}

func readFromChannel(namesChan chan string,wg *sync.WaitGroup) {
	defer wg.Done()
	for name := range namesChan {
		fmt.Println(name)
	}
}