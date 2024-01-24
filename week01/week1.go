package main

import "fmt"

func DeletedByIndex[T any](data []T, index int) ([]T, error) {
	length := len(data)
	if index < 0 || index > length {
		return nil, fmt.Errorf("下标超出范围，长度 %d, 下标 %d", length, index)
	}
	length--
	for i := index; i < length; i++ {
		data[i] = data[i+1]
	}
	data = data[:length]
	return Shrink(data), nil
}

func Shrink[T any](data []T) []T {
	c, l := cap(data), len(data)
	n, change := calCapacity(c, l)
	if !change {
		return data
	}
	nowData := make([]T, 0, n)
	nowData = append(nowData, data...)
	return nowData
}

func calCapacity(c, l int) (int, bool) {
	if c < 64 {
		return c, false
	}

	if c > 2048 && (c/l) >= 2 {
		return int(float32(c) * float32(0.625)), true
	}
	if c < 2048 && (c/l) >= 4 {
		return c / 2, true
	}
	return c, false
}
