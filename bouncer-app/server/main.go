package main

import ("fmt"
        "github.com/bouncer-app/proto"
        "google.golang.org/grpc"
        "encoding/json"
        "github.com/bouncer-app/database"
        "net"
        "context"
        "log"
      )

type server struct{
  db *database.Database
}

func NewServer()(*server, error){
  s := &server{}
  var err error
  s.db, err = database.OpenDB()
  if err != nil {
    log.Printf("Error opening Database Err:%s", err.Error())
    return nil, err
  }
  return s, nil
}

func main () {
  s, err := NewServer()
  if err != nil {
     panic(err)
  }
  listener, err := net.Listen("tcp", ":4040")
  if err != nil {
    panic(err)
  }

  srv := grpc.NewServer()
  proto.RegisterAddServiceServer(srv, s)

  if e := srv.Serve(listener); e != nil {
    panic(e)
  }
}

func (s *server) CreateUser(ctx context.Context, request *proto.CreateUserRequest) (*proto.CreateUserResponse, error){
  FN := request.GetFirstName()
  LN := request.GetLastName()
  AD := request.GetAddress()

  var user database.User

  user.FirstName = FN
  user.LastName = LN
  user.Address = AD

  err := s.db.CreateUser(user)
  if err != nil {
    panic(err)
  }
  if err == nil {
    fmt.Println("User Created Succesfully")
  }
  return &proto.CreateUserResponse{}, nil
}

  func (s *server) RetrieveUser(ctx context.Context, request *proto.GetRequest) (*proto.GetResponse, error) {
  key := request.GetID()
  data, err := s.db.RetrieveUser(key)
  if err != nil{
    panic(err)
  }
  if data == nil {
    fmt.Println("No data found")
    return nil,err
  }
  var user database.User

  err = json.Unmarshal(data, &user)
  if err != nil {
    panic(err)
  }
  uid := &user.ID
  id := uid.String()
  return &proto.GetResponse{ID : id, FirstName:user.FirstName, LastName:user.LastName, Address:user.Address}, nil
}

func (s *server) UpdateUser(ctx context.Context, request *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error){
  key := request.GetID()
  FN := request.GetFirstName()
  LN := request.GetLastName()
  AD := request.GetAddress()

  user := database.User{FirstName:FN, LastName:LN, Address:AD}

  err := s.db.UpdateUser(key, user)
  if err == nil{
    fmt.Println("User Data Updated Successfuly")
  } else {
    panic(err)
  }

  return &proto.UpdateUserResponse{}, nil
}

func (s *server) DeleteUser(ctx context.Context, request *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error){
  key := request.GetID()
  err := s.db.DeleteUser(key)
  if err != nil {
    panic(err)
  }

  return &proto.DeleteUserResponse{}, nil
}


func (s *server) RetrieveAllUsers(ctx context.Context, request *proto.GetAllRequest) (*proto.GetAllResponse, error) {
  data, err := s.db.ListUsers()
  if data == nil {
    log.Fatalln("No Data Found.")
  }
  if err != nil{
    panic(err)
  }
  var user database.User
  var pUser []*proto.GetResponse
  for _, li := range data{
      err = json.Unmarshal(li, &user)
      if err != nil {
        panic(err)
      }
      uid := user.ID
      key := uid.String()
      pUser = append(pUser, &proto.GetResponse{ID:key, FirstName:user.FirstName, LastName:user.LastName, Address:user.Address})
  }
  return &proto.GetAllResponse{User:pUser}, nil
}
