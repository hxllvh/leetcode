func twoSum(nums []int, target int) []int {
    search := map[int]int{}
    for i := range nums {
        i2, exists := search[target-nums[i]]
        if exists {
            return []int{i2, i}
        } else {
            search[nums[i]] = i
        }
    }
    return nil
}
