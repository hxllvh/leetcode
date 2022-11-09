func twoSum(nums []int, target int) []int {
    for i, val := range nums {
        for i2, val2 := range nums[i+1:] {
            if val + val2 == target {
                return []int{i, i2+i+1}
            }
        }
    }
    return nil
}
