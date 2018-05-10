1. What would a function signature that took a slice of numbers and adds them together look like?
  * sumSlice(xSlice []int) (int total) {}

2. Write a function that takes an integer and halves it and returns true if it was even or false if it was odd. For example, half(1) should return (0, false) and half(2) should return(1, true)
  * [./exercises/ex02.go](https://github.com/JackBurdick/learning_go/blob/master/introducing_go/chapter_06/exercises/ex02.go)

3. Write a function with one variadic parameter than finds the greatest number in a list of numbers.
  * [./exercises/ex03.go](https://github.com/JackBurdick/learning_go/blob/master/introducing_go/chapter_06/exercises/ex03.go)

4. Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers
  * [./exercises/ex04.go](https://github.com/JackBurdick/learning_go/blob/master/introducing_go/chapter_06/exercises/ex04.go)

5. The Fibonacci seq. is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function that can find fib(n).
  * [./exercises/ex05.go](https://github.com/JackBurdick/learning_go/blob/master/introducing_go/chapter_06/exercises/ex05.go)

6. What are defer, panic, and recover?  How do you recover from a runtime panic?
```
  * defer
    * called within a func, is executed just before the function is returned
  * panic
    * causes func to terminate (plus any func that call it --> program)
  * recover
    * prevents panic from going all the way up/down the call stack
    * [./example/9_panicRecover.go](https://github.com/JackBurdick/learning_go/blob/master/introducing_go/chapter_06/example/9_panicRecover.go)
```
7.  How do you get the memory address of a variable?
  * using the & operator
    * address of var => `&var`

8.  How do you assign a value to a pointer?
  * using the * operator
    * `*var = val`

9.  How do you create a new pointer?
  * using the `new` function
    * `newVal = new(int)`

10.  What is the value of x after running this program?
```
  func square(x *float64) {
      *x = *x * *x
  }
  func main() {
      x := 1.5
      square(&x)
  }
```
  * [./exercises/ex10.go](https://github.com/JackBurdick/learning_go/blob/master/introducing_go/chapter_06/exercises/ex10.go)
    * works as expected --> 1.5^2 = 2.25

11. Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).
  * [./exercises/ex11.go](https://github.com/JackBurdick/learning_go/blob/master/introducing_go/chapter_06/exercises/ex11.go)
