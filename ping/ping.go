package ping

func SenderPing(c chan string, done chan struct{}) {
	for  {
		c <- "ping"
		<-done
	}
}