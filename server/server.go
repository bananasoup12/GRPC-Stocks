
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
    "time"
    
    

	"google.golang.org/grpc"

	pb "github.com/GRPC-Stocks/routeguide"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)

var stockDB map[string]*pb.Stock

type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer


	mu         sync.Mutex // protects routeNotes
	routeNotes map[string][]*pb.RouteNote
}

// GetStock returns a stock of given name
func (s *routeGuideServer) GetStock(ctx context.Context, stockname *pb.StockName) (*pb.Stock, error) {
    
    stockName := stockname.Name
    if _, ok := stockDB[stockName]; ok {
        return stockDB[stockName], nil
    }

	// No stock was found, return nil
	return nil, nil
}

// Creates a new stock
func (s *routeGuideServer) CreateStock(ctx context.Context, stockupdate *pb.StockUpdate) (*pb.Error, error) {
    
    currentTime := time.Now()
    stockName := stockupdate.Name
    historicalpriceinfo := []*pb.HistoricalPriceInfo{}
    info1 := &pb.HistoricalPriceInfo{Date: currentTime.String(), Price: stockupdate.Price}
    historicalpriceinfo = append(historicalpriceinfo, info1)

    if _, ok := stockDB[stockName]; ok {
        return &pb.Error{Code: 1, Info: "Stock already exists"}, nil
    } else{
        stockDB[stockName] = &pb.Stock{Name: stockName, Historicalinfo: historicalpriceinfo}
    }

	
	return &pb.Error{Code: 0, Info: "No Error"}, nil
}

// Updates an existing stock by appending the new price to the historical price info array.
func (s *routeGuideServer) UpdateStock(ctx context.Context, stockupdate *pb.StockUpdate) (*pb.Error, error) {
    
    currentTime := time.Now()
    stockName := stockupdate.Name
    historicalpriceinfo := []*pb.HistoricalPriceInfo{}
    info1 := &pb.HistoricalPriceInfo{Date: currentTime.String(), Price: stockupdate.Price}
    historicalpriceinfo = append(historicalpriceinfo, info1)

    if _, ok := stockDB[stockName]; ok {
        historicalpriceinfo = stockDB[stockName].Historicalinfo
        historicalpriceinfo = append(historicalpriceinfo, info1)

        stockDB[stockName] = &pb.Stock{Name: stockName, Historicalinfo: historicalpriceinfo}
    } else{
        return &pb.Error{Code: 2, Info: "Stock does not exist"}, nil
    }

	
	return &pb.Error{Code: 0, Info: "No Error"}, nil
}



func newServer() *routeGuideServer {
	s := &routeGuideServer{routeNotes: make(map[string][]*pb.RouteNote)}
	
	return s
}


func initDB(){
    stockDB = make(map[string]*pb.Stock)
}

func main() {

    initDB()
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

