
//     | N1 | N2 | N3 | N4 | N5
// ----|------------------------
// R1  | x  | y  | x  | x  |  y
// ----|
// R2  | .. | .. | .. | .. | ..
// ----|
// R3  | .. | .. | .. | .. | ..
// ----|
// R4  | .. | .. | .. | .. | ..
// ----|
// R5  | .. | .. | .. | .. | ..
// ```
// [
// 	{ "R1":{
// 			 "x" : ["N1","N3", "N4"],
// 			 "y" : ["N2", "N5"],  //Order within the array does not matter
// 			 "z" : ["N5"],
// 			 // .. similarly for any other values present except `null`
// 		  }	,
// 	   "R2":{
// 	   	....
// 	   },
// 	   ...
// 	}
// ]
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

//Output of both the functions
//***********First function************
// [
//     {
//         "R1": {
//             "x": [
//                 "N1",
//                 "N3",
//                 "N4"
//             ],
//             "y": [
//                 "N2",
//                 "N5"
//             ]
//         },
//         "R2": {
//             "z": [
//                 "N1",
//                 "N5"
//             ],
//             "x": [
//                 "N2",
//                 "N4"
//             ],
//             "y": [
//                 "N3"
//             ]
//         },
//         "R3": {
//             "y": [
//                 "N1",
//                 "N4"
//             ],
//             "x": [
//                 "N2",
//                 "N3"
//             ],
//             "z": [
//                 "N5"
//             ]
//         },
//         "R4": {
//             "x": [
//                 "N1"
//             ],
//             "z": [
//                 "N2",
//                 "N3",
//                 "N5"
//             ],
//             "y": [
//                 "N4"
//             ]
//         },
//         "R5": {
//             "x": [
//                 "N1",
//                 "N4"
//             ],
//             "y": [
//                 "N2",
//                 "N3"
//             ],
//             "z": [
//                 "N5"
//             ]
//         }
//     }
// ]


//***********Second function************
// [
//     {
//         "R1": {
//             "x": [
//                 "N1",
//                 "N3",
//                 "N4"
//             ],
//             "y": [
//                 "N2",
//                 "N5"
//             ]
//         }
//     },
//     {
//         "R2": {
//             "z": [
//                 "N1",
//                 "N5"
//             ],
//             "x": [
//                 "N2",
//                 "N4"
//             ],
//             "y": [
//                 "N3"
//             ]
//         }
//     },
//     {
//         "R3": {
//             "y": [
//                 "N1",
//                 "N4"
//             ],
//             "x": [
//                 "N2",
//                 "N3"
//             ],
//             "z": [
//                 "N5"
//             ]
//         }
//     },
//     {
//         "R4": {
//             "x": [
//                 "N1"
//             ],
//             "z": [
//                 "N2",
//                 "N3",
//                 "N5"
//             ],
//             "y": [
//                 "N4"
//             ]
//         }
//     },
//     {
//         "R5": {
//             "x": [
//                 "N1",
//                 "N4"
//             ],
//             "y": [
//                 "N2",
//                 "N3"
//             ],
//             "z": [
//                 "N5"
//             ]
//         }
//     }
// ]