package main

import "fmt"

func playingDomino(card [][]int,  deck []int) interface{}{
	for i := 0 ; i < len(card); i++{
		if card[i][0] == deck[0] || card[i][0] == deck[1]{
			return card[i]
		}
	}
	return "tutup kartu"
}

func main(){
	fmt.Println(playingDomino([][]int{[]int{6,5},[]int{3,4},[]int{2,1},[]int{3,3}},[]int{4,3}))
	fmt.Println(playingDomino([][]int{[]int{6,5},[]int{3,3},[]int{3,4},[]int{2,1}},[]int{3,6}))
	fmt.Println(playingDomino([][]int{[]int{6,6},[]int{2,4},[]int{3,6}},[]int{5,1}))
}