
let matrix = [["x", "y", "x", "x", "y"],
["z", "x", "y", "x", "z"],
["y", "x", "x", "y", "z"],
["x", "z", "z", "y", "z"],
["x", "y", "y", "x", "z"]];

//***********Solution 1***************/
//Transform the matrix into required format
function dataMapper() {
    let result = [];
    let innerObj = {}
    for (let r = 0; r < matrix.length; r++) {
        let variable = "R" + (r + 1)
        innerObj[variable] = []
        innerObj[variable] = innerMapper(matrix[r])
    }
    result.push(innerObj)
    return result
}
//Transform the array into inner object like { "x": ["N1", "N2" ...]}
function innerMapper(matrixInnerArray) {
    let innerObj = {}
    for (let n = 0; n < matrixInnerArray.length; n++) {
        if (!innerObj[matrixInnerArray[n]]) {
            innerObj[matrixInnerArray[n]] = []
        }
        innerObj[matrixInnerArray[n]].push("N" + (n + 1))
    }
    return innerObj;
}

//***********Solution 2***************/
//Map the matrix into required format
function usingMapAndReduce() {
    return matrix.map((elm, i) => {
        return {
            ["R" + (i + 1)]: reduceFunction(elm)
        }
    })
}

//reduce the inner array into format { "x": ["N1", "N2" ...]}
function reduceFunction(innderArr) {
    return innderArr.reduce((prev, current, i) => {
        if (!prev[current]) {
            prev[current] = []
        }
        prev[current].push("N" + (i + 1))
        return { ...prev }
    }, {})
}


console.log(JSON.stringify(dataMapper(), 0, 1))
console.log(JSON.stringify(usingMapAndReduce(), 0, 1))