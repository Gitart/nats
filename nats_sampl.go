package main

import (
    "fmt"
    "log"
    "time"
    "github.com/nats-io/nats.go"
)



func main(){
    Clss()

}



func Clss(){

fmt.Println(nats.DefaultURL)    
// Connect to a server
nc, err := nats.Connect(nats.DefaultURL)

// Disable reconnect attempts
// nc, err := nats.Connect("demo.nats.io", nats.NoReconnect())
if err != nil {
   log.Fatal(err)
}

defer nc.Close()

// nc.Subscribe("foo", func(m *nats.Msg) {
//     fmt.Printf("Received a message: %s\n", string(m.Data))
// })

// nc.Subscribe("fooo", func(m *nats.Msg) {
//     fmt.Printf("Received a message: %s\n", string(m.Data))
// })



for {
     tm:=time.Now().Format("15:04:05")
    // ss:=`Pub = ` + tm 
    // Simple Publisher
    // nc.Publish("foo", []byte("Hello World"))
   
    nc.Publish("foo",     []byte("{'title':'ddd', 'time': " + tm + "}"))
    nc.Publish("fooo",    []byte("Foo3    -- " + tm))
    nc.Publish("testing", []byte("Testing -- " + tm))
    nc.Publish("t.com",   []byte("T.COM -- " + tm))


    fmt.Println("Send... ")
    time.Sleep(time.Second * 3)
}


fmt.Println("Ok")
// // Simple Async Subscriber
// nc.Subscribe("foo", func(m *nats.Msg) {
//     fmt.Printf("Received a message: %s\n", string(m.Data))
// })



}


// func Clss3(){
// // Connect to a server
// nc, _ := nats.Connect(nats.DefaultURL)

// // Simple Publisher
// nc.Publish("foo", []byte("Hello World"))

// // Simple Async Subscriber
// nc.Subscribe("foo", func(m *nats.Msg) {
//     fmt.Printf("Received a message: %s\n", string(m.Data))
// })

// // Responding to a request message
// nc.Subscribe("request", func(m *nats.Msg) {
//     m.Respond([]byte("answer is 42"))
// })

// // Simple Sync Subscriber
// sub, err := nc.SubscribeSync("foo")
// m, err   := sub.NextMsg(timeout)

// // Channel Subscriber
// ch := make(chan *nats.Msg, 64)
// sub, err := nc.ChanSubscribe("foo", ch)
// msg := <- ch

// // Unsubscribe
// sub.Unsubscribe()

// // Drain
// sub.Drain()

// // Requests
// msg, err := nc.Request("help", []byte("help me"), 10*time.Millisecond)

// // Replies
// nc.Subscribe("help", func(m *nats.Msg) {
//     nc.Publish(m.Reply, []byte("I can help!"))
// })

// // Drain connection (Preferred for responders)
// // Close() not needed if this is called.
// nc.Drain()

// // Close connection
// nc.Close()
// }
