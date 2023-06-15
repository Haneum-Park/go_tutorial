package channels

func ChannelBased() {
	ch := make(chan int)

	go func() {
		ch <- 123 // ? 채널에 123을 보냄
	}()

	ii := <-ch // ? 채널로부터 123을 받음
	println(ii)
}
