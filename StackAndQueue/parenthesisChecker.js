function parenthesisChecker(s) {
    let stack = s.split("")
    let temStack = []
    let found = true
    for (let i = 0; i < stack.length; i++) {
        switch (stack[i]) {
            case "[":
            case "(":
            case "{":
                temStack.push(stack[i])
                break;
            case "]": {
                let elem = temStack.pop()
                if (elem != "[") {
                    found = false
                }
                break;
            }
            case ")": {
                let elem = temStack.pop()
                if (elem != "(") {
                    found = false
                }
                break;
            }
            case "}": {
                let elem = temStack.pop()
                if (elem != "{") {
                    found = false
                }
                break;
            }
            default:
                break;
        }
        if (!found) {
            return false
        }
    }
    if (temStack.length != 0) {
        return false
    }
    return true
}

console.log(parenthesisChecker("()[]{}"))