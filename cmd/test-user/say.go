package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	pb "text2speech/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	output := flag.String("o", "test-result/output.wav", "wav file where the outpute will be saved")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Printf("usage:\n\t%s \"text to speack\"", flag.Arg(0))
		os.Exit(1)
	}

	conn, err := grpc.Dial(*backend, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to %s: %v", *backend, err)
	}

	defer conn.Close()
	// fmt.Printf("conn: %v\n", conn)
	client := pb.NewTextToSpeechClient(conn)
	text := &pb.Text{Text: flag.Arg(0)}
	// fmt.Printf("backend %s, arg1 %s\n", *backend, flag.Arg(0))

	res, err := client.Say(context.Background(), text)
	if err != nil {
		log.Fatalf("could not say %q: %v", text.Text, err)
	}

	// Create the output directory if it doesn't exist
	dir := filepath.Dir(*output)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("could not create directory %s: %v", dir, err)
	}
	if err := os.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("could not save %s: %v", *output, err)
	}

}
