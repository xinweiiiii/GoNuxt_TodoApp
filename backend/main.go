// package main


// // https://github.com/gorilla/mux.git
// import ("fmt"; "log"; "net/http"; "encoding/json"; "github.com/gorilla/mux"; "io/ioutil"; "go.mongodb.org/mongo-driver/bson"; "go.mongodb.org/mongo-driver/mongo"; "go.mongodb.org/mongo-driver/mongo/options"; "go.mongodb.org/mongo-driver/mongo/readpref"; "context"; "time")

// func homePage(w http.ResponseWriter, r *http.Request){
//     fmt.Fprintf(w, "Welcome to the HomePage!")
//     fmt.Println("Endpoint Hit: homePage")
// }

// func returnAllTasks(w http.ResponseWriter, r *http.Request){
//     fmt.Println("Endpoint Hit: Return All Tasks")
//     // Encoding the json string and write it as response
//     collection := client.Database("tasks").Collection("tasks")
//     var task Task
//     err := collection.FindOne(context.TODO(), filter).Decode(&task)
//     if err != nil {
//         log.Fatal(err)   
//     }
//     // json.NewEncoder(w).Encode(Tasks)
    
// }


// func returnSingleTasks(w http.ResponseWriter, r *http.Request){
//     vars := mux.Vars(r)
//     key := vars["id"]

//     // Extract the key from the URL and return it 
//     fmt.Fprintf(w, "Key: " + key)

//     for _, task := range Tasks {
//         if task.Id == key {
//             json.NewEncoder(w).Encode(task)
//         }
//     }
// }

// func createNewTask(w http.ResponseWriter, r *http.Request){
//     reqBody, _ := ioutil.ReadAll(r.Body)
//     var tasks Task
//     json.Unmarshal(reqBody, &tasks)
//     collection := client.Database("tasks").Collection("tasks")
//     insertResult, err := collection.InsertOne(context.TODO(), tasks)
//     if err != nil { 
//         log.Fatal(err)
//     }
//     // Tasks = append(Tasks, tasks)
//     // json.NewEncoder(w).Encode(tasks)
// }

// func deleteTask(w http.ResponseWriter, r *http.Request){
//     vars := mux.Vars(r)
//     id := vars["id"]

//     for index, Task := range Tasks {
//         if Task.Id == id {
//             Tasks = append(Tasks[:index], Tasks[index+1:]...)
//         }
//     }
// }


// func handleRequests() {
//     myRouter := mux.NewRouter().StrictSlash(true)
//     myRouter.HandleFunc("/", homePage)
//     myRouter.HandleFunc("/all", returnAllTasks)
//     myRouter.HandleFunc("/task/{id}", returnSingleTasks)
//     myRouter.HandleFunc("/task", createNewTask).Methods("POST")
//     myRouter.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
//     log.Fatal(http.ListenAndServe(":10000", myRouter))
// }

// func Connect() {
//     // Database Config
//     clientOptions := options.Client().ApplyURI("mongodb+srv://golang:gogo>@cluster0.uu0jq.mongodb.net/test")
//     client, err := mongo.NewClient(clientOptions)

//     //Set up a context required by mongo.Connect
//     ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//     err = client.Connect(ctx)

//     //Cancel context to avoid memory leak
//     defer cancel()
    
//     // Ping our db connection
//     err = client.Ping(context.Background(), readpref.Primary())
//     if err != nil {
//         log.Fatal("Couldn't connect to the database", err)
//     } else {
//         log.Println("Connected!")
//     }

//     // Connect to the database
//     db := client.Database("tasks")
//     return
// }

// type Task struct {
//     Id string `json:"Id"`
// 	Title string `json:"Title"`
//     Desc string `json:"desc"`
//     Priority string `json:"priority"`
//     Date string `json:"Data"`
// }

// var Tasks []Task

// func main() {

//     Tasks = []Task{
//         Task{Id: "1", Title: "SMU", Desc: "Finish up Iot Medium Article", Priority: "Medium", Date: "Thursday 26 August 2020"},
//         Task{Id: "2", Title: "Personal", Desc: "Apply Internship", Priority: "Medium", Date: "Thursday 26 August 2020"},
//     }
//     handleRequests()
// }


package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/cavdy-play/go_mongo/config"
	"github.com/cavdy-play/go_mongo/routes"
)

func main()  {
	// Database
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}






