function atoi(str) {
    // return Math.max(Math.min(parseInt(str) || 0, 2147483647), -2147483648)
    return str.charCodeAt(4) - 48
}

console.log(atoi("4193 with words"))