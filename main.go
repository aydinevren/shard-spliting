package main

import "fmt"

func main() {
	simulate(100, 4)
}

func simulate(size int, shard_size int) {
	mod_size := size % shard_size
	new_size := size - mod_size
	part := int(new_size / shard_size)

	for i := 0; i <= new_size-int(part); i += int(part) {
		var j = 0
		if i == new_size-int(part) {
			j = size - 1
		} else {
			j = i + int(part) - 1
		}
		fmt.Println(i, j)
	}
}
