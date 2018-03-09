package bina

func BinarySearch(list []int, target int) bool {
  halfPoint := len(list) / 2
  middleElement := list[halfPoint]
  if middleElement == target {
    return true
  } else if len(list) != 1 {
    if middleElement > target {
      return BinarySearch(list[:halfPoint], target)
    } else {
      return BinarySearch(list[halfPoint:], target)
    }
  }
  return false
}
