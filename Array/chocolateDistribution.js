function sort() {
    let arr = [12, 4, 7, 9, 2, 23,
        25, 41, 30, 40, 28,
        42, 30, 44, 48, 43,
        50]
    arr = arr.sort(function (a, b) { return a - b });
    console.log(arr)
    //need to find the block of size n and take the diffrence of min and max
}
sort()