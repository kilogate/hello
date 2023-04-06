package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/apaxa-go/helper/strconvh"
	"github.com/samber/lo"
)

type Stu struct {
	Name string
}

func (s Stu) Clone() Stu {
	return Stu{Name: s.Name}
}

func main() {
	// Filter：过滤需要的元素
	filter := lo.Filter([]string{"a", "b", "c", "123", "321", "A", "B", "C"}, func(item string, index int) bool {
		if len(item) == 3 || index == 0 {
			return true
		}
		return false
	})
	log.Printf("%v\n", filter) //  [a 123 321]

	// Map：映射元素
	mp := lo.Map([]int{11, 22, 33}, func(item int, index int) string {
		return fmt.Sprintf("%d%d", index, item)
	})
	log.Printf("%v\n", mp) // [011 122 233]

	// FilterMap：过滤并映射
	filterMap := lo.FilterMap([]string{"a", "b", "c", "123", "321", "A", "B", "C"}, func(item string, index int) (string, bool) {
		if len(item) == 3 || index == 0 {
			return fmt.Sprintf("%d%s", index, item), true
		}
		return "", false
	})
	log.Printf("%v\n", filterMap) // [0a 3123 4321]

	// FlatMap：映射并打平
	flatMap := lo.FlatMap([]string{"abc", "def"}, func(item string, index int) []string {
		return strings.Split(item, "")
	})
	log.Printf("%v\n", flatMap) // [a b c d e f]

	// Reduce & ReduceRight：化简
	reduce := lo.Reduce([]int{1, 2, 3, 4}, func(agg int, item int, index int) int {
		return agg + item
	}, 100)
	log.Printf("%v\n", reduce) // 110
	reduceRight := lo.ReduceRight([]int{1, 2, 3, 4}, func(agg int, item int, index int) int {
		return agg + item
	}, 100)
	log.Printf("%v\n", reduceRight) // 110

	// ForEach：循环
	lo.ForEach([]int{1, 2, 3}, func(item int, index int) {
		fmt.Println(index, item)
	})

	// Times
	result := lo.Times(3, func(i int) string {
		return strconvh.FormatInt(i)
	})
	log.Printf("%v\n", result) //  [0 1 2]

	// Uniq：去重
	uniq := lo.Uniq([]int{1, 3, 5, 6, 7, 3, 1})
	log.Printf("%v\n", uniq) // [1 3 5 6 7]

	// UniqBy：去重
	uniqBy := lo.UniqBy([]int{1, 2, 3, 4, 5, 2, 4, 6, 8}, func(item int) string {
		return strconvh.FormatInt(item % 2)
	})
	log.Printf("%v\n", uniqBy) // [1 2]

	// GroupBy：分组
	groupBy := lo.GroupBy([]int{1, 2, 3, 4, 5, 2, 4, 6, 8}, func(item int) string {
		return strconvh.FormatInt(item % 2)
	})
	log.Printf("%v\n", groupBy) //  map[0:[2 4 2 4 6 8] 1:[1 3 5]]

	// Chunk：分块
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	chunk := lo.Chunk(ints, 4)
	ints[0] = 999             // 结果共用底层数组
	log.Printf("%v\n", chunk) // [[999 2 3 4] [5 6 7 8] [9 10 11 12] [13 14 15 16] [17 18 19 20]]

	// PartitionBy：GroupBy Values 的有序版本，返回的是切片而不是映射
	partitionBy := lo.PartitionBy([]int{1, 2, 3, 4, 5, 2, 4, 6, 8}, func(item int) string {
		return strconvh.FormatInt(item % 2)
	})
	log.Printf("%v\n", partitionBy) // [[1 3 5] [2 4 2 4 6 8]]

	// Flatten：打平
	flatten := lo.Flatten([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	log.Printf("%v\n", flatten) // [1 2 3 4 5 6 7 8 9]

	// Interleave：交替
	interleave := lo.Interleave[int]([]int{1}, []int{2, 5, 8}, []int{3, 6}, []int{4, 7, 9, 10})
	log.Printf("%v\n", interleave) // [1 2 3 4 5 6 7 8 9 10]

	// Shuffle：乱序
	shuffle := lo.Shuffle([]int{1, 3, 5, 2, 4, 6})
	log.Printf("%v\n", shuffle) // [6 1 3 5 4 2]

	// Reverse：反序
	reverse := lo.Reverse([]int{1, 3, 5, 7, 9})
	log.Printf("%v\n", reverse) // [9 7 5 3 1]

	// Fill：填充
	fill := lo.Fill([]Stu{{"A"}, 7: {"G"}}, Stu{"F"})
	log.Printf("%v\n", fill) // [{F} {F} {F} {F} {F} {F} {F} {F}]

	// Repeat：重复
	repeat := lo.Repeat(8, Stu{"F"})
	log.Printf("%v\n", repeat) // [{F} {F} {F} {F} {F} {F} {F} {F}]

	// RepeatBy
	repeatBy := lo.RepeatBy(5, func(i int) string {
		return strconvh.FormatInt(i)
	})
	log.Printf("%v\n", repeatBy) //  [0 1 2 3 4]

	// KeyBy：[]V -> map[K]V
	keyBy := lo.KeyBy([]*Stu{{"A"}, {"B"}, {"C"}}, func(item *Stu) string {
		return item.Name
	})
	log.Printf("%v\n", keyBy) // map[A:0x140001046a0 B:0x140001046b0 C:0x140001046c0]

	// Associate == SliceToMap：[]T -> map[K]V
	associate := lo.Associate([]*Stu{{"A"}, {"B"}, {"C"}}, func(item *Stu) (string, string) {
		return item.Name, item.Name
	})
	log.Printf("%v\n", associate) //  map[A:A B:B C:C]

	// SliceToMap == Associate：[]T -> map[K]V
	sliceToMap := lo.SliceToMap([]*Stu{{"A"}, {"B"}, {"C"}}, func(item *Stu) (string, string) {
		return item.Name, item.Name
	})
	log.Printf("%v\n", sliceToMap) //  map[A:A B:B C:C]

	// Drop
	drop := lo.Drop([]int{1, 2, 3, 4, 5, 6, 7}, 5)
	log.Printf("%v\n", drop) // [6 7]

	// DropRight
	dropRight := lo.DropRight([]int{1, 2, 3, 4, 5, 6, 7}, 5)
	log.Printf("%v\n", dropRight) // [1 2]

	// DropWhile
	dropWhile := lo.DropWhile([]int{1, 2, 3, 4, 5, 6, 7}, func(item int) bool {
		return item != 5
	})
	log.Printf("%v\n", dropWhile) // [5 6 7]

	// DropRightWhile
	dropRightWhile := lo.DropRightWhile([]int{1, 2, 3, 4, 5, 6, 7}, func(item int) bool {
		return item != 5
	})
	log.Printf("%v\n", dropRightWhile) // [1 2 3 4 5]

	// Reject
	reject := lo.Reject([]string{"a", "b", "c", "123", "321", "A", "B", "C"}, func(item string, index int) bool {
		if len(item) == 3 || index == 0 {
			return true
		}
		return false
	})
	log.Printf("%v\n", reject) // [b c A B C]

	// Count：元素数量
	count := lo.Count([]string{"a", "b", "c", "123", "321", "A", "B", "C"}, "A")
	log.Printf("%d\n", count) // 1

	// CountBy
	countBy := lo.CountBy([]string{"a", "b", "c", "123", "321", "A", "B", "C"}, func(item string) bool {
		return item == "A"
	})
	log.Printf("%d\n", countBy) // 1

	// CountValues：Value的数量
	countValues := lo.CountValues([]string{"A", "B", "C", "123", "321", "A", "B", "C"})
	log.Printf("%v\n", countValues) //  map[123:1 321:1 A:2 B:2 C:2]

	// CountValuesBy
	countValuesBy := lo.CountValuesBy([]string{"A", "B", "C", "123", "321", "A", "B", "C"}, func(item string) int {
		return len(item)
	})
	log.Printf("%v\n", countValuesBy) // map[1:6 3:2]

	// Subset：子集
	subset := lo.Subset([]int{1, 2, 3, 4, 5, 6, 7}, 2, 3)
	log.Printf("%v\n", subset) // [3 4 5]

	// Slice：切片
	slice := lo.Slice([]int{1, 2, 3, 4, 5, 6, 7}, 2, 8)
	log.Printf("%v\n", slice) // [3 4 5 6 7]

	// Replace：替换
	replace := lo.Replace([]int{1, 2, 3, 4, 5}, 1, 10, -1)
	log.Printf("%v\n", replace) // [10 2 3 4 5]

	// ReplaceAll
	replaceAll := lo.ReplaceAll([]int{1, 2, 3, 4, 5}, 1, 10)
	log.Printf("%v\n", replaceAll) // [10 2 3 4 5]

	// Compact
	compact := lo.Compact([]int{0, 1, 2, 3, 4, 0})
	log.Printf("%v\n", compact) // [1 2 3 4]

	// IsSorted：是否排序
	isSorted := lo.IsSorted([]int{0, 1, 2, 3, 4, 6, 5})
	log.Printf("%t\n", isSorted) // false

	// IsSortedByKey
	isSortedByKey := lo.IsSortedByKey([]int{0, 1, 2, 3, 4, 6, 5}, func(item int) int {
		return item
	})
	log.Printf("%t\n", isSortedByKey) // false
}
