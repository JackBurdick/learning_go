## Questions

```
// send "ping" to channel c
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// print message received from channel c
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
```

* is ping constantly "sending"/"attempting to send" messages in the channel?
  * inefficiency?
