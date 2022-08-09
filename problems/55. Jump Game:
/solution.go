func canJump(nums []int) bool {
  var reach bool[len(nums)]
  reach[len(nums) - 1] = true

  for i := len(nums) - 1; i >= 0; i-- {
    r := false
    for j := 0; j < nums[i]; j++ {
      if nums[i + j] {
        r = true
        break
      }
    }
    reach[i] = r
  }

  return reach[0]
}
