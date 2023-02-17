package main

import(
	"encoding/json"
	"log"
	"strings"
	"net/http"
	"time"
	"context"
	"os"
	"time"
	"os/signal"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"//mongo db bilan aloqa qilib beradi
	"gopkg.in/mgo.v2/bson"

)
var rnd *renderer.Render
var db *mgo.Database

const(
	hostName 		string = "localhost:27017"
	dbName 			string = "demo_todo"
	colectionName   string = "todo"
	port 			string = ":9000"
)

type(
	todoModel struct{// maluotlar bazasi bilang ishlagani kerak
		ID 			bson.ObjectId `bson:"_id,omitempty"`
		Title       string `bson:"title"`
		Completed   bool `bson:"completed"`
		CreatedAt   time.Time `bson:"createdAt`
	}
	todo struct{
		ID 		string `json:"id"`
		Titel   string `json:"titel"`
		Completes string `json:"completed"`
		CreatedAt time.Time `jsom:"created_at"`
	}	
)
func init(){
	rnd = renderer.New()
	sess, err := mgo.Dial(hostName)//mongo db bilan bog'lab beradi
	checkErr(err)
	sess.SetMode(mgo.Monotonic, true)

}
