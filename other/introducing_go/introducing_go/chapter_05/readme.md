1. How do you access the fourth element of an array or slice?
  * zero based, arr[3]

2. What is the length of the slice created using make([]int, 3, 9)?
  * make(type, length, capacity) --> 3 is the length

3. Given the following array, what would x[2:5] give you?
  ```
    x := [6]string{"a","b","c","d","e","f"}
  ```
  * [./exercises/ex3.go](https://github.com/JackBurdick/learning_go/blob/master/introducing_go/chapter_05/exercises/ex3.go)
    * `[2:5)` --> `"c","d","e"`

4. Write a program that finds the smallest number in this list;
  ```
    x := []int{
      49,96,86,68,
      57,82,63,70,
      37,34,83,27,
      19,97, 9,17,
    }
  ```
  * [./exercises/ex4.go](https://github.com/JackBurdick/learning_go/blob/master/introducing_go/chapter_05/exercises/ex4.go)
