function equilibriumPoint(arr) {
    let end = arr.length - 1
    let i = 0
    let leftSum = arr[i]
    i++
    let rightSum = arr[end]
    end--
    let result = 0
    while (i <= end) {
        if (leftSum < rightSum) {
            leftSum = leftSum + arr[i]
            i++
        } else if (rightSum < leftSum) {
            rightSum = rightSum + arr[end]
            end--
        } else {
            if (i == end) {
                result = i
                i++
            }
        }
    }
    return i
}
console.log(equilibriumPoint([1,3,5,2,2]))