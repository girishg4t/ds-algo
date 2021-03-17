## Data Trasformar

### Problem statment:
Given an array in below format
```
    | N1 | N2 | N3 | N4 | N5
----|------------------------
R1  | x  | y  | x  | x  |  y
----|
R2  | .. | .. | .. | .. | ..
----|
R3  | .. | .. | .. | .. | ..
----|
R4  | .. | .. | .. | .. | ..
----|
R5  | .. | .. | .. | .. | ..
```
transform it into below output
```
[
	{ "R1":{
			 "x" : ["N1","N3", "N4"],
			 "y" : ["N2", "N5"],  //Order within the array does not matter
			 "z" : ["N5"],
			 // .. similarly for any other values present except `null`
		  }	,
	   "R2":{
	   	....
	   },
	   ...
	}
]
```

### Solution:
How to Run 
``` 
node index.js
```
Input array:
```
[["x", "y", "x", "x", "y"],  
["z", "x", "y", "x", "z"],  
["y", "x", "x", "y", "z"],  
["x", "z", "z", "y", "z"],  
["x", "y", "y", "x", "z"]]
```
***********First function output************
```
[
    {
        "R1": {
            "x": [
                "N1",
                "N3",
                "N4"
            ],
            "y": [
                "N2",
                "N5"
            ]
        },
        "R2": {
            "z": [
                "N1",
                "N5"
            ],
            "x": [
                "N2",
                "N4"
            ],
            "y": [
                "N3"
            ]
        },
        "R3": {
            "y": [
                "N1",
                "N4"
            ],
            "x": [
                "N2",
                "N3"
            ],
            "z": [
                "N5"
            ]
        },
        "R4": {
            "x": [
                "N1"
            ],
            "z": [
                "N2",
                "N3",
                "N5"
            ],
            "y": [
                "N4"
            ]
        },
        "R5": {
            "x": [
                "N1",
                "N4"
            ],
            "y": [
                "N2",
                "N3"
            ],
            "z": [
                "N5"
            ]
        }
    }
]
```

***********Second function output************
```
[
    {
        "R1": {
            "x": [
                "N1",
                "N3",
                "N4"
            ],
            "y": [
                "N2",
                "N5"
            ]
        }
    },
    {
        "R2": {
            "z": [
                "N1",
                "N5"
            ],
            "x": [
                "N2",
                "N4"
            ],
            "y": [
                "N3"
            ]
        }
    },
    {
        "R3": {
            "y": [
                "N1",
                "N4"
            ],
            "x": [
                "N2",
                "N3"
            ],
            "z": [
                "N5"
            ]
        }
    },
    {
        "R4": {
            "x": [
                "N1"
            ],
            "z": [
                "N2",
                "N3",
                "N5"
            ],
            "y": [
                "N4"
            ]
        }
    },
    {
        "R5": {
            "x": [
                "N1",
                "N4"
            ],
            "y": [
                "N2",
                "N3"
            ],
            "z": [
                "N5"
            ]
        }
    }
]
```