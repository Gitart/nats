
package main


// import (
//     "fmt"
//     "github.com/google/uuid"
// )

// func main() {
//     fmt.Println(uuid.Must(uuid.NewRandom()))
// }

import (
    "fmt"
    "time"
    "github.com/nats-io/nats.go"
    // "lib" 
)


var GlobalNats struct {
    Nats *nats.Conn
    Error error
}


// Ini file - Yaml
// CLI 

// Init 
func init(){

     // fmt.Println(lib.Unixtime())

     GlobalNats.Nats, GlobalNats.Error = nats.Connect(nats.DefaultURL)
     if GlobalNats.Error != nil {
        fmt.Println("Error")
        return
     }
}



// Main
func main(){
    

    var topic = "bar"

    fmt.Println("Subscribe...", nats.DefaultURL)
    // nc, err := nats.Connect(nats.DefaultURL)
    // if err!=nil {
    //    fmt.Println(err.Error())  
    // }
    // defer nc.Close()

 // nts:=GlobalNats.Nats
 //    nts.Subscribe(topic, func(m *nats.Msg) {
 //          // fmt.Printf("Received request: %s, final stop, sending back to %v\n", m.Data, m.Reply)
 //          fmt.Printf("%s\n", m.Data)
 //          // nc.Flush()
 //    })


     Subscribe(topic)             
    
     Pub(topic, "Test : ")
    
}


func Subscribe(topic string) {

    fmt.Println("Subscribe...")
     // fmt.Println("Subscribe...", nats.DefaultURL)
    // nc, err := nats.Connect(nats.DefaultURL)
    // if err!=nil {
    //    fmt.Println(err.Error())  
    // }
    // defer nc.Close()
    nts:=GlobalNats.Nats

    nts.Subscribe(topic, func(m *nats.Msg) {
          // fmt.Printf("Received request: %s, final stop, sending back to %v\n", m.Data, m.Reply)
          
          // fmt.Printf("%s\n", nc.Status())
          fmt.Printf("%s\n", m.Data)
          nts.Flush()
    })
}

// func Publish(topic, txt string){

//     for {
//         tm:=time.Now().Format("150405")
//         Pub(topic, txt + " # " + tm)        
//         time.Sleep( 1000 * time.Millisecond)
//     }
// }



// Publish 
func Pub(topic, txt string){
     nts := GlobalNats.Nats

     // nc, err := nats.Connect(nats.DefaultURL)
     // defer nc.Close()

     // if err!=nil {
     //    fmt.Println(err.Error())  
     // }


    for {
       
       tm:=time.Now().Format("150405")
       
       // Simple Publisher
       nts.Publish(topic, []byte(txt+"-"+tm))

       // fmt.Printf("%s\n", nc.Status())
       // nc.Flush()
       time.Sleep( 1000 * time.Millisecond)
    }

}

