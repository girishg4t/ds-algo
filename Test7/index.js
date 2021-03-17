//Given a string, check if it is a palindrome by ignoring spaces. E.g. race car would be a palindrome.

//Approch
//take whole as a string
//start utill middle
//take first and last element and check if it is equal 
//if space comes ignore

function palindrome(str) {
    let count = str.length - 1;
    let first = str[0]
    let last = str[count]
    for (let i = 1; i < Math.floor(str.length / 2); i++) {
        if (first == " ") {
            first = i + 1;
        }
        if (last == " ") {
            last = count - 1;
        }
        if (str[first] != str[last]) {
            return false;
        }
        count--;
    }
    return true;
}
//console.log(palindrome("Too hot to hoot"))
console.log(palindrome("Abc def ghi jklm"))
