package main

import "fmt"

// question : https://quera.ir/problemset/contest/34081/%D8%B3%D8%A4%D8%A7%D9%84-%D8%B1%DB%8C%D8%A7%D8%B6%DB%8C%D8%A7%D8%AA-%D8%AF%D8%A7%DB%8C%D8%B1%D9%87-%D8%B9%D8%AC%DB%8C%D8%A8
func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	count := 1
	for it := 1 + k; it%n != 1; it += k {
		count++
	}

	fmt.Println(count)

}
