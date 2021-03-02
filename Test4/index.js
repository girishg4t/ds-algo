//Given an array A, divide it into two arrays say B and C, such that the sum of elements in array B is greater than or
// equal to the sum of elements in array C and array B should have a minimum number of elements.

//var A = [1, 2, 3, 4, 5, 6]
//var B = [5, 6]
//var c = [1, 2, 3, 4]

//Approch
//1) it's batter to sort the array
//2) copy all elements in c array
//3) start coping last element of c array into b utill sum of all the values in array b is greater then c

function calculateArray(arr) {
    var c = [...arr];
    var b = [];
    for (let i = c.length - 1; i > -1; i--) {
        b.push(c[i]);
        if (getSum(b) > getSum(c.slice(0, i))) {
            c = c.slice(0, i);
            break;
        }
    }
    console.log(b);
    console.log(c);
}
function getSum(arr) {
    return arr.reduce(function (accumulator, current) {
        return accumulator + current;
    });
}
//calculateArray([1, 2, 3, 4, 5, 6])

//Split an array into two equal Sum subarrays

function equalArray(arr) {
    var c = [...arr];
    var b = [];
    for (let i = c.length - 1; i > -1; i--) {
        b.push(c[i]);
        if (getSum(b) == getSum(c.slice(0, i))) {
            c = c.slice(0, i);
            break;
        }
    }
    console.log(b);
    console.log(c);
}
equalArray([4, 1, 2, 3 ])