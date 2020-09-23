package main

import "fmt"

func main() {
	ids := []int{23, 65, 15, 6, 2, 43}

	// loop through ids
	for i, id := range ids {
		fmt.Printf("%d -> %d\n", i, id)
	}

	// loop through ids without indices
	// put an underscore instead of index
	for _, id := range ids {
		fmt.Printf("%d\n", id)
	}

	// sum values using range
	sum := 0
	for _, id := range ids {
		sum += id
	}
	println("Sum", sum)

	// range with map
	people := map[string]string{"Mehdi": "Teymorian", "Alex": "Johnas"}
	for key, value := range people {
		fmt.Printf("%s: %s\n", key, value)
	}

	// only keys of map
	for key := range people {
		println("Name:" + key)
	}
}
