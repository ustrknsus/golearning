// Arrays, slices, and maps
// Exercise: Count how many times each word occurs in a slice.
package main

import "fmt"

func main() {
	fixed := [3]int{1, 2, 3}
	hosts := []string{"server-a", "server-b"}
	hosts = append(hosts, "server-c")
	ports := map[string]int{"ssh": 22, "https": 443}
	fmt.Println(fixed)
	for index, host := range hosts {
		fmt.Println(index, host)
	}
	if port, ok := ports["ssh"]; ok {
		fmt.Println("SSH port:", port)
	}
}
