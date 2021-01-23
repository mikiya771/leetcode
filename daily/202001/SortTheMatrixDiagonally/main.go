package main

import (
	"sort"
)

func diagonalSort(mat [][]int) [][]int {
	h := len(mat)
	if h == 0 {
		return [][]int{}
	}
	w := len(mat[0])
	enc := mapEncoder(mat, w, h)
	for _, l := range enc {
		sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })
	}
	return mapDecoder(enc, w, h)
}

//通常のmapからDiagonal にしたものに変換する
func mapEncoder(mat [][]int, w, h int) [][]int {
	ret := [][]int{}
	for i := 0; i < w+h; i++ {
		tmp := []int{}
		for j := 0; j <= i; j++ {
			if j < h && w-1+j-i >= 0 {
				tmp = append(tmp, mat[j][w-1+j-i])
			}
		}
		ret = append(ret, tmp)
	}
	return ret
}

//Diagonal にしたものから通常のmapに変換する
func mapDecoder(mat [][]int, w, h int) [][]int {
	ret := [][]int{}
	for j := 0; j < h; j++ {
		tmp := []int{}
		for i := 0; i < w; i++ {
			tmp = append(tmp, 0)
		}
		ret = append(ret, tmp)
	}
	for k, l := range mat {
		for m, e := range l {
			if w-k-1 > 0 {
				ret[m][w-k+m-1] = e
			} else {
				ret[m+1+(k-w)][m] = e
			}
		}
	}
	return ret
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
