package main

import ("context"
        "google.golang.org/grpc"
        //"net/http"
        //"github.com/satori/go.uuid"
        "github.com/bouncer-app/proto"
        //"github.com/bouncer-app/database"
        //"encoding/json"
        "log"
        "time"
       )

func main() {
  conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
  if err != nil {
    panic(err)
  }
  client := proto.NewAddServiceClient(conn)

/*                  Create User                                      */
  /*request := &proto.CreateUserRequest{FirstName:"Raviraj", LastName:"Shinde", Address:"Italy"}

  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()

  _, err = client.CreateUser(ctx, request)
  if err != nil{
    log.Fatalln(err)
  } else {
    log.Println("User Created")
  }*/

  /*                  Retrieve User                                       */
  /*  request := &proto.GetRequest{ID:"1c276a84-438e-484e-b483-1dc98803d57c"}

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

   response, err := client.RetrieveUser(ctx, request)
   if err != nil {
     log.Fatalln(err)
   }

  log.Println(response)*/

  /*                                Update User                                 */
  /*request := &proto.UpdateUserRequest{ID:"71e5fb2b-0b8e-4600-a5aa-b91d22d6412f", FirstName:"Raviraj", LastName:"Shinde", Address:"Chandigarh"}

  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()

  _, err = client.UpdateUser(ctx, request)
  if err != nil{
    log.Fatalln(err)
  } else {
    log.Println("User Updated")
  }*/
/*                           Delete User                                      */
/*request := &proto.DeleteUserRequest{ID:"83dee1c3-0b31-4886-8655-9f651e88934f"}

ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()

_, err = client.DeleteUser(ctx, request)
if err != nil{
  log.Fatalln(err)
} else {
  log.Println("User Deleted")
}*/


/*                  List Users                                       */
/*  request := &proto.GetAllRequest{}

  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()

  response, err := client.RetrieveAllUsers(ctx, request)
  if err != nil{
    log.Fatalln(err)
  }
  log.Fatalln(response)*/



}
