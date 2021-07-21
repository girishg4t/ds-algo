/**
 * Initialize your data structure here.
 */
 var MyQueue = function() {
   this.stack1 = []
   this.stack2 = [] 
};

/**
 * Push element x to the back of queue. 
 * @param {number} x
 * @return {void}
 */
MyQueue.prototype.push = function(x) {
    this.stack1.push(x)
    while(this.stack1.length != 0){
        let el = this.stack1.pop()
        this.stack2.push(el)
    }
};

/**
 * Removes the element from in front of queue and returns that element.
 * @return {number}
 */
MyQueue.prototype.pop = function() {
    let el = this.stack2.pop()
    return el
};

/**
 * Get the front element.
 * @return {number}
 */
MyQueue.prototype.peek = function() {
    let el = this.stack2.pop()

    return el
};

/**
 * Returns whether the queue is empty.
 * @return {boolean}
 */
MyQueue.prototype.empty = function() {
    return this.stack2.length == 0 ? true : false
};

/** 
 * Your MyQueue object will be instantiated and called as such:
 * var obj = new MyQueue()
 * obj.push(x)
 * var param_2 = obj.pop()
 * var param_3 = obj.peek()
 * var param_4 = obj.empty()
 */

//[1,2,3,4,5]- stack
//[]