package main

import (
	"fmt"
	"sort"
)

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func appendNum(slice []int) {
	slice = append(slice, 4)
}

func appendNumPointer(slice *[]int) {
	*slice = append(*slice, 4)
}

func appendNumNewSlice(slice []int) []int {
	slice = append(slice, 4)
	return slice
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5} 
	
	changeArray(array) // array를 array2로 값(40byte) 복사(C or C++은 pointer로 변환) 
	changeSlice(slice) // slice는 array의 data field 값, 사용량(length), capacity 정보를 담은 pointer(24byte structure)
	
	fmt.Println(array) // [1 2 3 4 5]
	fmt.Println(slice) // [1 2 200 4 5]
	
	appendNum(slice) // len, cap이 모두 5이므로 복사하여 새로운 slice에 append
	fmt.Println(slice) // [1 2 200 4 5]
	
	appendNumPointer(&slice) // pointer(8byte) copy
	fmt.Println(slice) // [1 2 200 4 5 4]
	
	slice = appendNumNewSlice(slice) // slice structure(24byte) 복사(pointer와 성능 차이 크게 없음 / golang 지향)
	fmt.Println(slice) // [1 2 200 4 5 4 4]
	
	fmt.Println(slice, len(slice), cap(slice)) // [1 2 200 4 5 4 4] 7 10
	// cap(=10) = 초기 cap(=5) * 2 (부족할 때, 배수 단위로 증가)
	
	slice2:= slice[1:3:5] // slice[시작 index : 끝 index : 최대 index]
	fmt.Println(slice2, len(slice2), cap(slice2)) // [2 200] 2 4
	
	
	// append를 이용한 copy	
	slice3 := []int{5, 4, 3, 2, 1}
	slice4 := append([]int{}, slice3...) // slice 복사(빈 slice에 slice3의 전체 인자를 복사)
	
	slice3[1] = 100
	fmt.Println(slice3) // [5 100 3 2 1]
	fmt.Println(slice4) // [5 4 3 2 1]
	
	
	// 내장함수 copy 이용	
	slice5 := []int{9, 8, 7, 6, 5}
	slice6 := make([]int, len(slice5))
	copy(slice6, slice5) // slice 복사 copy(dst, src)
	
	slice5[1] = 100
	fmt.Println(slice5) // [9, 100, 7, 6, 5]
	fmt.Println(slice6) // [9, 8, 7, 6, 5]
	
	
	// slice 요소 삭제
	slice7 := []int{1 ,2, 3, 4, 5, 6}
	idx := 2 // 지우고 싶은 index
	slice7 = append(slice7[:idx], slice7[idx+1:]...)
	fmt.Println(slice7) // [1, 2, 4, 5, 6]
	
	
	// slice 요소 삽입
	slice8 := []int{1 ,2, 3, 4, 5, 6}
	idx2 := 2 // 삽입하고 싶은 index
	slice8 = append(slice8[:idx2], append([]int{100}, slice8[idx2:]...)...) // 불필요한 메모리 할당(temporary buffer)
	fmt.Println(slice8) // [1, 2, 100, 3, 4, 5, 6]
	
	
	// copy 사용하여 slice 요소 삽입(불필요한 메모리 할당이 없음)
	slice9 := []int{1 ,2, 3, 4, 5, 6}
	idx3 := 2 // 삽입하고 싶은 index
	slice9 = append(slice9, 0)
	copy(slice9[idx3+1:], slice9[idx3:])
	slice9[idx] = 100
	fmt.Println(slice9) // [1, 2, 100, 3, 4, 5, 6]
	
	
	// sort
	sort.Ints(slice9)
	fmt.Println(slice9) // [1, 2, 3, 4, 5, 6, 100]
		
}
