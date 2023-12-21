package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"

	pb "text2speech/protobuf"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port to listen on")
	flag.Parse()

	log.Infof("listening on port %d", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("could not listen to port %d: %v", *port, err)
	}

	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}
}

type server struct{}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {
	f, err := os.CreateTemp("", "say")
	if err != nil {
		return nil, fmt.Errorf("could not create temp file: %v", err)
	}
	log.Infof("Processing Request for: %s", text.Text)
	if err := f.Close(); err != nil {
		log.Errorf("could not close temp file %s: %v", f.Name(), err)
		return nil, fmt.Errorf("could not close temp file %s: %v", f.Name(), err)
	}

	cmd := exec.Command("flite", "-t", text.Text, "-o", f.Name())
	if data, err := cmd.CombinedOutput(); err != nil {
		log.Errorf("could not run flite for input %s: %v", data, err)
		return nil, fmt.Errorf("could not run flite: %s", data)
	}

	data, err := os.ReadFile(f.Name())
	if err != nil {
		log.Errorf("could not read temp file %s: %v", f.Name(), err)
		return nil, fmt.Errorf("could not read temp file %s: %v", f.Name(), err)
	}
	log.Infof("Returning Response for: %s", text.Text)
	return &pb.Speech{Audio: data}, nil
}
