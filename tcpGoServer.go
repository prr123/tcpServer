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
	log.Printf("listening on %s\n", listenAddr)

	for {
		con, err := l.Accept()
		if err != nil {log.Fatalf("Accept: %v\n", err)}

		PrintCon(&con)

		go handleRequest(con)
	}

}

	func PrintCon(con *net.Conn) {

    loc := (*con).LocalAddr()
    fmt.Printf("Src Adr: %s %s\n", loc.Network(), loc.String())
    rem :=(*con).RemoteAddr()
    fmt.Printf("Dest Adr: %s %s\n", rem.Network(), rem.String())

}


func handleRequest(conn net.Conn) {
	// incoming request
	// notice we are creating a buffer for each request!
	buf := make([]byte, 1024)
	bufLen, err := conn.Read(buf)
	if err != nil {log.Fatalf("con.Read: %v\n",err)}

	idx := bytes.Index(buf, []byte("STOP"))
	if idx > -1 {
		log.Println("Received STOP -- Exiting TCP server!")
		os.Exit(1)
	}

	fmt.Print("s-> ", string(buf[:bufLen]))
	// write data to response
	timStr := time.Now().Format(time.RFC1123)
	respStr := fmt.Sprintf("Your message is: %s. Received time: %s", string(buf[:bufLen]), timStr)
	conn.Write([]byte(respStr))

	// close conn
	conn.Close()
}
