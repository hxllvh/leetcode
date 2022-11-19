func isPalindrome(x int) bool {
    str := strings.Split(strconv.Itoa(x), "")
    for i:=0; i < len(str); i++ {
        if str[len(str)-1-i] != str[i] {
            return false
        }
    }
    return true
}
