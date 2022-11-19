func isPalindrome(x int) bool {
    var number int
    n := x
    for n > 0 {
        number = number*10 + n%10
        n = n/10
    }
    return number == x
}
