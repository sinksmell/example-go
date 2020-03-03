package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
)

var (
	mockCtx, cancel = context.WithCancel(context.Background())
)

func main() {
	l, err := net.Listen("tcp", ":23333")
	if err != nil {
		panic(err)
	}

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("pprof failed %v", err)
		}
	}()

	conns := make([]net.Conn, 0)
	defer func() {
		cancel()
		for _, conn := range conns {
			conn.Close()
		}
	}()

	for {

		select {
		case <-mockCtx.Done():
			fmt.Println("context is down")
		default:
			conn, err := l.Accept()
			if err != nil {
				if ne, ok := err.(net.Error); ok && ne.Temporary() {
					log.Printf("accept temp err %v\n", err)
					continue
				}
				log.Printf("accept err %v\n", err)
				return
			}

			go handleConn(mockCtx, conn)
			conns = append(conns, conn)
			if len(conns)%100 == 0 {
				log.Printf("total number of conns: %v", len(conns))
			}
		}
	}

}

func handleConn(ctx context.Context, conn net.Conn) {
	ctx = context.WithValue(ctx, "test", "test")
	resp := http.Response{
		Status:     http.StatusText(200),
		StatusCode: 200,
		Proto:      "HTTP/1.1",
		ProtoMajor: 0,
		ProtoMinor: 0,
		Header:     http.Header{},
		Body:       nil,
	}
	bytes, _ := json.Marshal(resp)
	//io.Copy(os.Stdout, conn)
	fmt.Fprintf(conn, "%v %v %v\r\n", resp.Proto, resp.StatusCode, resp.Status)
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(bytes))
	fmt.Fprintf(conn, "Content-Type: application/json\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, "%s\r\n", string(bytes))
}
