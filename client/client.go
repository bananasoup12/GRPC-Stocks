
package main

import (
	"context"
	"flag"
	
	"log"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
	pb "github.com/GRPC-Stocks/routeguide"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name used to verify the hostname returned by the TLS handshake")
)

//Print info for a given stock
func printStock(client pb.RouteGuideClient, stockname *pb.StockName) {
	log.Printf("Getting Stock with name (%d)", stockname.Name)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stock, err := client.GetStock(ctx, stockname)
	if err != nil {
		log.Fatalf("%v.GetStocks(_) = _, %v: ", client, err)
	}
	log.Println(stock)
}

//Create a new stock
func createStock(client pb.RouteGuideClient, stockupdate *pb.StockUpdate) {
	log.Printf("Create Stock with info (%d, %d)", stockupdate.Name, stockupdate.Price)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	error, err := client.CreateStock(ctx, stockupdate)
	if err != nil {
		log.Fatalf("%v.CreateStocks(_) = _, %v: ", client, err)
	}
	log.Println(error)
}

//Update an existing stock 
func updateStock(client pb.RouteGuideClient, stockupdate *pb.StockUpdate) {
	log.Printf("Update Stock with info (%d, %d)", stockupdate.Name, stockupdate.Price)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	error, err := client.UpdateStock(ctx, stockupdate)
	if err != nil {
		log.Fatalf("%v.UpdateStocks(_) = _, %v: ", client, err)
	}
	log.Println(error)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRouteGuideClient(conn)
	
	separator := strings.Repeat("--", 50)

	createStock(client, &pb.StockUpdate{Name: "Gamestop", Price: 4})
	createStock(client, &pb.StockUpdate{Name: "AMC", Price: 1})
	log.Println(separator)
	time.Sleep(1 * time.Second)

	updateStock(client, &pb.StockUpdate{Name: "Gamestop", Price: 5})
	updateStock(client, &pb.StockUpdate{Name: "AMC", Price: 2})
	printStock(client, &pb.StockName{Name: "Gamestop"})
	printStock(client, &pb.StockName{Name: "AMC"})
	log.Println(separator)
	time.Sleep(1 * time.Second)

	updateStock(client, &pb.StockUpdate{Name: "Gamestop", Price: 10})
	updateStock(client, &pb.StockUpdate{Name: "AMC", Price: 5})
	printStock(client, &pb.StockName{Name: "Gamestop"})
	printStock(client, &pb.StockName{Name: "AMC"})
	log.Println(separator)
	time.Sleep(1 * time.Second)

	updateStock(client, &pb.StockUpdate{Name: "Gamestop", Price: 20})
	updateStock(client, &pb.StockUpdate{Name: "AMC", Price: 8})
	printStock(client, &pb.StockName{Name: "Gamestop"})
	printStock(client, &pb.StockName{Name: "AMC"})
	log.Println(separator)
	time.Sleep(1 * time.Second)

	updateStock(client, &pb.StockUpdate{Name: "Gamestop", Price: 30})
	updateStock(client, &pb.StockUpdate{Name: "AMC", Price: 12})
	printStock(client, &pb.StockName{Name: "Gamestop"})
	printStock(client, &pb.StockName{Name: "AMC"})
	log.Println(separator)
	time.Sleep(1 * time.Second)

	updateStock(client, &pb.StockUpdate{Name: "Gamestop", Price: 50})
	updateStock(client, &pb.StockUpdate{Name: "AMC", Price: 16})
	printStock(client, &pb.StockName{Name: "Gamestop"})
	printStock(client, &pb.StockName{Name: "AMC"})
	log.Println(separator)
	time.Sleep(1 * time.Second)

	updateStock(client, &pb.StockUpdate{Name: "Gamestop", Price: 100})
	updateStock(client, &pb.StockUpdate{Name: "AMC", Price: 18})
	printStock(client, &pb.StockName{Name: "Gamestop"})
	printStock(client, &pb.StockName{Name: "AMC"})
	log.Println(separator)
	time.Sleep(1 * time.Second)

	updateStock(client, &pb.StockUpdate{Name: "Gamestop", Price: 200})
	updateStock(client, &pb.StockUpdate{Name: "AMC", Price: 25})
	printStock(client, &pb.StockName{Name: "Gamestop"})
	printStock(client, &pb.StockName{Name: "AMC"})
	log.Println(separator)


}
