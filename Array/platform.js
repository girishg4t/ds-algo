function platform() {
    let arr = [0900, 0940, 0950, 1100, 1500, 1800]
    let dep = [0910, 1120, 1130, 1200, 1900, 2000]
    let i = 0
    let j = 0
    let m = 0;
    let result = 0
    while (i < arr.length && j < dep.length) {
        if (arr[i] < dep[j]) {
            i++
            m++
        } else if (dep[j] <= arr[i]) {
            j++
            m--
        }
        if (m > result) {
            result = m
        }
    }
    console.log(result)
}
platform()
