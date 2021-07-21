function countInversion(){
    let arr = [2, 4, 1, 3, 5]
    mergeSort(arr, 0, arr.length-1)
}

function mergeSort(arr, i, j){
    if(i < j){
        n= Math.round(arr.length/2)
        mergeSort(arr)
    }
}