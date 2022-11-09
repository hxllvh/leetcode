func twoSum(nums []int, target int) []int {
    for i, val := range nums {
        for i2, val2 := range nums {
            if i2 <= i {
                continue
            } else if val + val2 == target {
                return []int{i, i2}
            }
        }
    }
    return nil
}
