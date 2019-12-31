
#### Constant Time: O(1)
An algorithm is said to run in constant time if it requires the same amount of time regardless of the input size. Examples:

array: accessing any element
fixed-size stack: push and pop methods
fixed-size queue: enqueue and dequeue methods

#### Linear Time: O(n)
An algorithm is said to run in linear time if its time execution is directly proportional to the input size, i.e. time grows linearly as input size increases. Examples:

array: linear search, traversing, find minimum
ArrayList: contains method
queue: contains method

#### Logarithmic Time: O(log n)
An algorithm is said to run in logarithmic time if its time execution is proportional to the logarithm of the input size. Example:

binary search
Recall the "twenty questions" game - the task is to guess the value of a hidden number in an interval. Each time you make a guess, you are told whether your guess iss too high or too low. Twenty questions game imploies a strategy that uses your guess number to halve the interval size. This is an example of the general problem-solving method known as binary search:

locate the element a in a sorted (in ascending order) array by first comparing a with the middle element and then (if they are not equal) dividing the array into two subarrays; if a is less than the middle element you repeat the whole procedure in the left subarray, otherwise - in the right subarray. The procedure repeats until a is found or subarray is a zero dimension.
Note, log(n) < n, when n→∞. Algorithms that run in O(log n) does not use the whole input.

merge sort

#### Quadratic Time: O(n2)
An algorithm is said to run in quadratic time if its time execution is proportional to the square of the input size. Examples:

bubble sort, selection sort, insertion sort, 

#### Algorithm's complexity - memory complexity vs cpu complexity
https://www.bigocheatsheet.com/


#### concurrency model of Golang, Communicating via sequential processing, go routines&channels which involves less context switching of the CPU as oppose to 
#### concurrency using shared memory such as mutex which involves heavy usage of the context switching hence it is slower