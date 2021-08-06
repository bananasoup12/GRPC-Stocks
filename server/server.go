/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a simple gRPC server that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It implements the route guide service whose definition can be found in routeguide/route_guide.proto.
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

// GetFeature returns the feature at the given point.
func (s *routeGuideServer) GetStock(ctx context.Context, stockname *pb.StockName) (*pb.Stock, error) {
    
    stockName := stockname.Name
    if _, ok := stockDB[stockName]; ok {
        return stockDB[stockName], nil
    }

	// No feature was found, return an unnamed feature
	return nil, nil
}

func (s *routeGuideServer) CreateStock(ctx context.Context, stockupdate *pb.StockUpdate) (*pb.Error, error) {
    
    currentTime := time.Now()
    stockName := stockupdate.Name
    historicalpriceinfo := []*pb.HistoricalPriceInfo{}
    info1 := &pb.HistoricalPriceInfo{Date: currentTime.String(), Price: 100.00}
    historicalpriceinfo = append(historicalpriceinfo, info1)

    if _, ok := stockDB[stockName]; ok {
        historicalpriceinfo = stockDB[stockName].Historicalinfo
        historicalpriceinfo = append(historicalpriceinfo, info1)

        stockDB[stockName] = &pb.Stock{Name: stockName, Historicalinfo: historicalpriceinfo}
    } else{
        stockDB[stockName] = &pb.Stock{Name: stockName, Historicalinfo: historicalpriceinfo}
    }

	// No feature was found, return an unnamed feature
	return &pb.Error{Code: 0, Info: "No Error"}, nil
}



func newServer() *routeGuideServer {
	s := &routeGuideServer{routeNotes: make(map[string][]*pb.RouteNote)}
	s.loadFeatures(*jsonDBFile)
	return s
}

// loadFeatures loads features from a JSON file.
func (s *routeGuideServer) loadFeatures(filePath string) {
	
	
	
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

