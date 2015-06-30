/*
Buffered Channels
It's also possible to pass a second parameter
to the make function when creating a channel:

c := make(chan int, 1)

This creates a buffered channel with a capacity of 1.
Normally channels are synchronous; both sides of the channel will wait
until the other side is ready. A buffered channel is asynchronous; sending or
receiving a message will not wait unless the channel is already full.
*/
