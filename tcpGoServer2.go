package main

import (
        "fmt"
        "net"
        "os"
		"log"
        "bytes"
        "time"
)

func main() {

	numArgs := len(os.Args)
	port:="10101"

	useStr := "tcpServer [port]"
	helpStr := fmt.Sprintf("help:\nThe program is a tcp server that listens on port: %s\n", port)
	if numArgs == 2 {
		if os.Args[1] == "help" {
			fmt.Printf("usage is: %s\n", useStr)
			fmt.Printf("%s\n", helpStr)
			os.Exit(1)
        }
		port = os.Args[1]
	}
    listenAddr := ":" + port
	log.Printf("listen address: %s\n", listenAddr)
	l, err := net.Listen("tcp", listenAddr)
	if err != nil { log.Fatalf("listen: %v\n", err)}
	defer l.Close()
	log.Printf("listening on %s start: %s\n", listenAddr, time.Now().Format(time.RFC1123))

	conCount :=0
	for {
		con, err := l.Accept()
		if err != nil {log.Fatalf("Accept: %v\n", err)}
		conCount++
		PrintCon(&con, conCount)
		go handleRequest(con, conCount)
	}

}

	func PrintCon(con *net.Conn, conId int) {

    loc := (*con).LocalAddr()
    fmt.Printf("%d: Src Adr: %s %s\n", conId, loc.Network(), loc.String())
    rem :=(*con).RemoteAddr()
    fmt.Printf("%d: Dest Adr: %s %s\n", conId, rem.Network(), rem.String())

}


func handleRequest(conn net.Conn, conId int) {
	// incoming request
	// notice we are creating a buffer for each request!
	buf := make([]byte, 1024)

	for {
		bufLen, err := conn.Read(buf)
		if err != nil {log.Fatalf("con.Read: %v\n",err)}

		idx := bytes.Index(buf, []byte("STOP"))
		if idx > -1 {
			log.Println("Received STOP -- Closing Connection!")
			conn.Close()
			return
		}

		fmt.Printf("%d:-> %s", conId, string(buf[:bufLen]))
	// write data to response
//	timStr := time.Now().Format(time.RFC1123)
//	respStr := fmt.Sprintf("%s", string(buf[:bufLen]))
		conn.Write(buf[:bufLen])
	}
	// close conn
}
