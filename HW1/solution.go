package main

import "fmt"

/*
Никогда не использовал generics и хотел попробовать.
Если есть какие-то best practice был бы рад узнать.
*/

type Number interface {
	~int | ~float64 | ~float32 | ~uint
}

func sortSlice[T Number](s []T) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	fmt.Printf("Sort Slice: %v\n", s)
}

func incrementOdd[T Number](slice []T) {
	for i := 0; i < len(slice); i++ {
		if i%2 == 0 {
			slice[i]++
		}
	}
	fmt.Printf("Odd incremented: %v\n", slice)
}

func printSlice[T Number](slice []T) {
	fmt.Printf("Print Slice: %v\n", slice)
}

func appendFunc[T Number](dst func([]T), src ...func([]T)) func([]T) {
	// не получилось без замыкания :((
	return func(slice []T) {
		dst(slice)
		for _, f := range src {
			f(slice)
		}
	}
}

func reverseSlice[T Number](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Printf("Reversed slice: %v\n", slice)
}
