package main

import "fmt"

/**
Bubble sort function created using knowledge I gained 
from `tou_of_go_1.md` in my go_notes section
*/
func BubbleSort(params []int) []int {
  noSwaps := false
  sorted := params
  for !noSwaps {
    numSwaps := 0
    for i := 1; i < len(params); i++ {
      curr := sorted[i]
      prev := sorted[i - 1]
      if (curr < prev) {
        sorted[i - 1] = curr;
        sorted[i] = prev
        numSwaps++
      }
    } 
    noSwaps = numSwaps == 0
  }
  return sorted
}


// It's somewhat recursive because it still has a for loop ;)
func BubbleSortSomewhatRecursive(params []int) []int {
  sorted := params  
  numSwaps := 0
    for i := 1; i < len(params); i++ {
      curr := sorted[i]
      prev := sorted[i - 1]
      if (curr < prev) {
        sorted[i - 1] = curr;
        sorted[i] = prev
        numSwaps++
      }
    } 
    
    if numSwaps > 0 {
      return BubbleSortSomewhatRecursive(sorted)
    } else {
      return sorted
    }

}

func main() {
  result1 := BubbleSort([]int{2, 6, 1, 2, 7, 4})
  result2 := BubbleSortSomewhatRecursive([]int{2, 6, 1, 2, 7, 4})
  fmt.Println("Iterative bubble sort results: ", result1)
  fmt.Println("Recursive bubble sort results", result2)
}
