package main

import (
	"fmt"
	"math/rand"
	"time"
)

func acakDadu() int {
	return rand.Intn(7-1) + 1
}

type Pemain struct {
	Poin int
	Dadu []int
}

func cekBerapaYangTidakKosong(temp map[int]Pemain) int {
	var tidakKosong int
	for _, pemain := range temp {
		if len(pemain.Dadu) > 0 {
			tidakKosong++
		}
	}
	return tidakKosong
}

func main() {
	n := 3
	m := 4

	var temp = make(map[int]Pemain)
	for i := 1; i <= n; i++ {
		temp[i] = Pemain{
			Dadu: []int{1},
		}
	}

	fmt.Println("========== Giliran 1")
	for i := 1; i <= n; i++ {
		var val = []int{}
		poin := temp[i].Poin
		arr := temp[i].Dadu
		for j := 1; j <= m; j++ {
			acak := acakDadu()
			val = append(val, acak)
			if acak == 6 {
				poin++
			} else if acak == 1 {
				a := i
				for {
					a++
					if a > n {
						a = 1
					}
					if len(temp[a].Dadu) > 0 {
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
				Dadu: arr[1:],
			}
		}
		fmt.Printf("Pemain %d %v \n", i, val)
		time.Sleep(time.Second / 3)
	}
	fmt.Println("========== Evaluasi 1")
	for i, pemain := range temp {
		fmt.Printf("Pemain %d (%d) %v \n", i, pemain.Poin, pemain.Dadu)
	}
	fmt.Println()

	b := 1
	for {
		b++
		if cekBerapaYangTidakKosong(temp) < 2 {
			break
		}
		fmt.Println("========== Giliran", b)
		for i, pemain := range temp {
			var val = []int{}
			if len(pemain.Dadu) > 0 {
				poin := temp[i].Poin
				temp[i] = Pemain{
					Poin: poin,
					Dadu: nil,
				}
				arr := temp[i].Dadu
				for k := 1; k <= len(pemain.Dadu); k++ {
					acak := acakDadu()
					val = append(val, acak)
					if acak == 6 {
						poin++
					} else if acak == 1 {
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
			fmt.Printf("Pemain %d %v \n", i, val)
			time.Sleep(time.Second / 3)
		}
		fmt.Println("========== Evaluasi", b)
		for i, pemain := range temp {
			fmt.Printf("Pemain %d (%d) %v \n", i, pemain.Poin, pemain.Dadu)
		}
		fmt.Println()

	}

	fmt.Println("=======================================")

	max := temp[1].Poin
	index := 1
	for i, v := range temp {
		fmt.Println(i, v)
		if v.Poin > max {
			max = v.Poin
			index = i
		}
	}

	fmt.Printf("Game dimenangkan oleh pemain nomor %d", index)

}
