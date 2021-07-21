function stockBuyandSell() {
    let arr = [4, 2, 2, 2, 4]
    let low = arr[0]
    let heigh = arr[0]
    for (let i = 1; i < arr.length; i++) {
        heigh = arr[i]
        if (low > heigh) {
            low = heigh
        } else {
            if (heigh > arr[i + 1]) {
                console.log(low + "=>" + heigh)
                i++
                low = arr[i]
                heigh = arr[i]
            }
        }
    }
    console.log(low + "=>" + heigh)
}
stockBuyandSell()