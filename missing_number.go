package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Print the missing number")
	seq := sequence()
	for x, s := range seq {
		fmt.Println("number [" + strconv.Itoa(x) + "] is " + strconv.Itoa(s))
	}
	//
	//lowest := lowest(seq)
	fmt.Println("+++++++++ Lowest number from enterd +++++++++++")
	fmt.Println("The lowest number of the sequence is " + strconv.Itoa(lowest(seq)))
	fmt.Println("========== Highest number from entered ========")
	fmt.Println("The highest number of the sequence is " + strconv.Itoa(highest(seq)))
	fmt.Println("******** An array with only unique values *****")
	fmt.Println(uniquev(seq))
	fmt.Println("---------- Array organised in asc order -------")
	fmt.Println(order(uniquev(seq)))
	fmt.Println("+++++ Missing number from sequence ++++++++++")
	fmt.Println(missing(uniquev(seq), len(uniquev(seq))))
}

//############################################################################################

// missing finds missing number
func missing(unique []int, n int) int {
	sum := 0
	nsum := 0

	for i := 0; i < n; i++ {
		sum += unique[i]
	}

	nsum = ((n + 1) * (n + 2)) / 2
	return (nsum - sum)
}

// order re-organises the array in asc order
func order(uniq []int) (sorted []int) {
	sort.Sort(sort.IntSlice(uniq))
	sorted = uniq
	return
}

// uniquev finds unique values of the sequence
func uniquev(seq []int) (unique []int) {
	for _, s := range seq {
		addv := true
		for _, u := range unique {
			if u == s {
				addv = false
				break
			} else {
				continue
			}
		}
		//
		if addv == true {
			unique = append(unique, s)
		}
	}
	return
}

// find highest number from the sequence
func highest(seq []int) (high int) {
	for x, s := range seq {
		if x == 0 {
			high = s
			continue
		}
		if high < s {
			high = s
		}
	}
	return
}

// find lowest positive number from sequence
func lowest(seq []int) (low int) {
	for x, s := range seq {
		if x == 0 {
			low = s
			continue
		}
		if low > s {
			low = s
		}
	}
	return
}

// Return an array of integers
func sequence() (seq []int) {
	var num int
	for {
		fmt.Println("Enter number > than 0 : ")
		fmt.Scanf("%v", &num)

		if num <= 0 {
			break
		} else {
			//fmt.Println("You entered " + strconv.Itoa(num) + "....")
			seq = append(seq, num)
		}
	}
	return
}
