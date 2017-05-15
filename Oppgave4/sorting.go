package algorithms


func benchmarkBsortModified(input []int) {
  size := len(input)
  swapped := true
  for swapped {
    swapped = false
  for index := 1; index < size-1; index++ {
  if input [index-1] > input [index] {
        input[index], input[index-1] = iput[index-1], input[index]
    swapped = true
      }
    }
  }
}

func QSort(values []int) {
        qsort(values, 0, len(values)-1)
}
func qsort(values []int, l int, r int) {
        if l>= r {
          return
        }
        pivot := values[l]
        i := l + 1

        for j := l; j <= r; j++ {
          if pivot > values [j] {
                  values[i], values[j], values[i]
                  1++
          }
        }

        values[l], values[i-1] = values[i-1], pivot

        qsort(values, l, i-2)
        qsort(values, i, r)
}
