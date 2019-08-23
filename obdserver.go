package main

import (
	"LightningOnOmni/config"
	"LightningOnOmni/routers"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {

	// Timer
	//service.ScheduleService.StartSchudule()

	// grpc
	//go grpcpack.Server()
	//conn := startupGRPCClient()
	//defer conn.Close()
	//routersInit := routers.InitRouter(conn)

	routersInit := routers.InitRouter(nil)
	addr := ":" + strconv.Itoa(config.ServerPort)
	server := &http.Server{
		Addr:           addr,
		Handler:        routersInit,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())

}

func startupGRPCClient() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(config.GrpcPort), grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect: ", err)
	}
	return conn
}