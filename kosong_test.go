package main

import (
	"fmt"
	"testing"
)

func TestKosong(t *testing.T) {
	var temp = map[int]Pemain{
		1: {
			Poin: 3,
			Dadu: []int{4, 1},
		},
		2: {
			Poin: 0,
			Dadu: []int{},
		},
		3: {
			Poin: 0,
			Dadu: []int{1, 5, 3, 2},
		},
	}

	for i, pemain := range temp {
		if len(pemain.Dadu) > 0 {
			poin := temp[i].Poin
			temp[i] = Pemain{
				Poin: poin,
				Dadu: nil,
			}
			arr := temp[i].Dadu
			for k := 0; k < len(pemain.Dadu); k++ {
				acak := acakDadu()
				fmt.Println("pemain", i, "dadu ", acak)
				if acak == 6 {
					poin++
				} else if acak == 1 {
					n := 4
					a := i + 1
					if a > n {
						a = 1
					}
					for {
						if len(temp[a].Dadu) == 0 {
							if a > n {
								a = 1
							} else {
								a++
							}
						}
						if len(temp[a].Dadu) != 0 {
							break
						}
					}
					poin2 := temp[a].Poin
					arr2 := temp[a].Dadu
					arr2 = append(arr2, acak)
					temp[a] = Pemain{
						Poin: poin2,
						Dadu: arr2,
					}
				} else {
					arr = append(arr, acak)
				}
				temp[i] = Pemain{
					Poin: poin,
					Dadu: arr,
				}
			}
		}
	}
	fmt.Println("=======================================")
	for i, v := range temp {
		fmt.Println(i, v)
	}
}

func TestHAHA(t *testing.T) {
	var a = []int{1}
	fmt.Println(len(a))
}
