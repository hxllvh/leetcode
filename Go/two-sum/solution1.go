func twoSum(nums []int, target int) []int {
    for i :=0; i < len(nums) - 1; i++ {
        for i2 := i + 1; i2 < len(nums); i2++ {
            if nums[i] + nums[i2] == target {
                return []int{i, i2}
            }
        }
    }
    return nil
}
