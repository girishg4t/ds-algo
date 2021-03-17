//Given an array, return the number of recurring elements.

//var a = [1,2,3,4,2,7,8,8,3];
//output = 2, 3, 8 

//Approch
//sort the array so it will be [1,2,2,3,3,4,7,8,8]
//now take first element and check with 2nd element unit not equal
//take now the other element

function getRecurringElements(arr) {
    var firstelem = arr[0];
    var finalArr = [];
    for (let i = 1; i < arr.length; i++) {
        if (firstelem == arr[i]) {
            finalArr.push(firstelem);
        }
        firstelem = arr[i];
    }
    return finalArr;
}

console.log(getRecurringElements([1,2,3,4,2,7,8,8,3].sort((a, b) => a - b)))