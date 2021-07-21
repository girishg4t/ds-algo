function rainWater() {
    let arr = [3, 0, 2, 0, 4]
    let start = 0
    let end = arr.length - 1
    let leftMaxHight = arr[start];
    let rightMaxHight = arr[end];
    let result = 0;
    while (start < end) {
        if (arr[start] <= arr[end]) {
            start++
            leftMaxHight = Math.max(leftMaxHight, arr[start])
            result += leftMaxHight - arr[start]
        } else {
            end--
            rightMaxHight = Math.max(rightMaxHight, arr[end])
            result += rightMaxHight - arr[end]
        }
    }
    console.log(result)
}


rainWater()