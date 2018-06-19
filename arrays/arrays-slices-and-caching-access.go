package main

func main() {

}

const ArraySize int = 65536

func GenerateArray() {
	var array = [ArraySize]int{}

	for i := 0; i < len(array); i++ {
		array[i] = i
	}
}

func GenerateSliceNoCapAppend() {
	var slice = []int{}

	for i := 0; i < ArraySize; i++ {
		slice = append(slice, i)
	}
}

func GenerateSliceWithCapAppend() {
	var slice = make([]int, 0, ArraySize)

	for i := 0; i < ArraySize; i++ {
		slice = append(slice, i)
	}
}

func GenerateSliceWithLenCap() {
	var slice = make([]int, ArraySize)

	for i := 0; i < ArraySize; i++ {
		slice[i] = i
	}
}

func GenerateArrayNoLen() {
	var array = [ArraySize]int{}

	for i := 0; i < ArraySize; i++ {
		array[i] = i
	}
}

func GenerateArrayCacheBewm() {
	var array = [ArraySize]int{}
	const gap = ArraySize / 32
	// 0 .. 1024 .. 2048 .. .. 1 .. 1025 .. 2049

	for i, j := 0, 0; i < ArraySize; i, j = i+1, i*gap%ArraySize+(i*gap/ArraySize) {
		// skip access
		array[j] = i
	}
}

// GenerateArrayNonCacheBewm
// serves as a control test vs cachebewm
// Uses same number of instructions as cachebewm, but writes sequentially
// thus using the cache normally.
func GenerateArrayCacheBewmControl() {
	var array = [ArraySize]int{}
	const gap = ArraySize / 32
	// 0 .. 1024 .. 2048 .. .. 1

	for i, j := 0, 0; i < ArraySize; i, j = i+1, i*gap%ArraySize+(i*gap/ArraySize) {
		// sequential access
		array[i] = j
	}
}

func GenerateArrayCacheBewm2() {
	var array = [ArraySize]int{}

	for i := 0; i < ArraySize; i++ {
		// flipping access
		if i%2 == 0 {
			array[i] = i
		} else {
			array[ArraySize-i] = i
		}
	}
}

func GenerateArrayCacheBewm2Control() {
	var array = [ArraySize]int{}

	for i := 0; i < ArraySize; i++ {
		// sequential access
		if i%2 == 0 {
			array[i] = i
		} else {
			array[i] = ArraySize - i
		}
	}
}
