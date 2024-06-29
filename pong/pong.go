package pong

func SenderPong(c chan string, done chan struct{}) {
	for {
		<- done
		c <- "pong"
	}
}