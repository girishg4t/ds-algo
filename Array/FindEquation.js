function FindEquation() {
    for (var a = 1; a < 100; a++) {
        for (var b = 1; b < 100; b++) {
            for (var c = 1; c < 100; c++) {
                for (var d = 1; d < 100; d++) {
                    if ((Math.pow(a, 3) + Math.pow(b, 3)) == (Math.pow(c, 3) + Math.pow(d, 3))) {
                        console.log(a + "," + b + "," + c + "," + d)
                    }
                }
            }
        }
    }
}

FindEquation()