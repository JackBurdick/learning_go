1. How do you specify the direction of a channel type?
    * using `<-`
        * Receiver
            * `<-chan int`
        * Send-only
            * `chan<- int`
2. Write your own Sleep function using time.After.
  ```
   func Sleep(duration time.Duration) {
        <-time.After(duration)
   }
  ```

3. What is a buffered channel? How would you create one with a capacity of 20?
  * buffered channel = channel that has a buffer
    * if the receiver is busy, the stores the sent message in a buffer
    * to make
        * `make(chan int, 20)`
