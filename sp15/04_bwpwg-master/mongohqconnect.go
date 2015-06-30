package main
 
import (
        "fmt"
        "labix.org/v2/mgo"
        "labix.org/v2/mgo/bson"
        "log"
        "os"
)
 
type Person struct {
        Name string
        Email string
}
 
func main() {
        // In the command window,
        // set MONGOHQ_URL=mongodb://IndianGuru:password@troup.mongohq.com:10080/godata
        // IndianGuru is my username, replace the same with yours
        uri := os.Getenv("MONGOHQ_URL")
        if uri == "" {
                fmt.Println("no connection string provided")
                os.Exit(1)
        }
 
        sess, err := mgo.Dial(uri)
        if err != nil {
                fmt.Printf("Can't connect to mongo, go error %v\n", err)
                os.Exit(1)
        }
        defer sess.Close()
        
        sess.SetSafe(&mgo.Safe{})
        
        collection := sess.DB("godata").C("user")

        err = collection.Insert(&Person{"Stefan Klaste", "klaste@posteo.de"},
	                        &Person{"Nishant Modak", "modak.nishant@gmail.com"},
	                        &Person{"Prathamesh Sonpatki", "csonpatki@gmail.com"},
	                        &Person{"murtuza kutub", "murtuzafirst@gmail.com"},
	                        &Person{"aniket joshi", "joshianiket22@gmail.com"},
	                        &Person{"Michael de Silva", "michael@mwdesilva.com"},
	                        &Person{"Alejandro Cespedes Vicente", "cesal_vizar@hotmail.com"})
        if err != nil {
                log.Fatal("Insert: ", err)
        }

        result := Person{}
        err = collection.Find(bson.M{"name": "Nishant Modak"}).One(&result)
        if err != nil {
                log.Fatal("Find: ", err)
        }

        fmt.Println("Email Id:", result.Email)
}
