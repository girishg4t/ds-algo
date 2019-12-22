function getRandomNumber(min, max) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

function getShuffedArray() {
    let cardArr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17,
        18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
        36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52];
    
    let numbersCopy = [...cardArr];
    let shuffledArray = [];
    for (let i = 0; i < cardArr.length; i++) {
        const random = getRandomNumber(0, 51 - i)
        shuffledArray[i] = numbersCopy[random];
        numbersCopy.splice(random, 1);

    }
    return shuffledArray;
}



console.log(JSON.stringify(getShuffedArray(), 0, 2));