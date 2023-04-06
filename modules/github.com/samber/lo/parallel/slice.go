package main

import (
	"fmt"
	"log"

	"github.com/apaxa-go/helper/strconvh"
	"github.com/samber/lo/parallel"
)

func main() {
	// Map
	mp := parallel.Map([]int{11, 22, 33}, func(item int, index int) string {
		return fmt.Sprintf("%d%d", index, item)
	})
	log.Printf("%v\n", mp) // [011 122 233]

	// ForEach
	parallel.ForEach([]int{1, 2, 3}, func(item int, index int) {
		fmt.Println(index, item)
	})

	// Times
	result := parallel.Times(3, func(i int) string {
		return strconvh.FormatInt(i)
	})
	log.Printf("%v\n", result) //  [0 1 2]

	// GroupBy
	groupBy := parallel.GroupBy([]int{1, 2, 3, 4, 5, 2, 4, 6, 8}, func(item int) string {
		return strconvh.FormatInt(item % 2)
	})
	log.Printf("%v\n", groupBy) //  map[0:[2 4 2 4 6 8] 1:[1 3 5]]

	// PartitionBy（GroupBy Values 的有序版本）
	partitionBy := parallel.PartitionBy([]int{1, 2, 3, 4, 5, 2, 4, 6, 8}, func(item int) string {
		return strconvh.FormatInt(item % 2)
	})
	log.Printf("%v\n", partitionBy) // [[1 3 5] [2 4 2 4 6 8]]
}
