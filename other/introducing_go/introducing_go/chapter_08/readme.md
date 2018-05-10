1. Why do we use packages?
  * Organization -- packages allow us to group together bits of relevant code

2. What is the difference between an identifier that starts with a capital letter and one that doesn't (Average vs average?)
  * In go, identifiers that start with a Capital letter are exported/visible, lower case identifiers are not

3. What is a package alias?  How do you make one?
  * similar to python's `import numpy as np`.. another name(presumably shorter) to identify a package
  * ` import f "fmt" `

4. Create Min and Max functions to find the minimum an maximum values in a slice of float64s.
  * [./exercises/ex04.go](https://github.com/JackBurdick/learning_go/blob/master/introducing_go/chapter_08/exercises/ex04.go)

5. How would you document the functions you created in 4?
  * [./exercises/ex04_documented.go](https://github.com/JackBurdick/learning_go/blob/master/introducing_go/chapter_08/exercises/ex04.go)
  * Documentation is created by adding a comment to the preceding line of a function
  * This documentation is visible in the godoc tool
