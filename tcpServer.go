package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
		"log"
        "strings"
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
    listenStr := ":" + port
	log.Printf("listen address: %s\n", listenStr)

	l, err := net.Listen("tcp", listenStr)
	if err != nil { log.Fatalf("listen: %v\n", err)}
	defer l.Close()
	log.Printf("listening on %s\n", listenStr)

	con, err := l.Accept()
	if err != nil {log.Fatalf("Accept: %v\n", err)}
	defer con.Close()
	PrintCon(&con)

	for {
		netData, err := bufio.NewReader(con).ReadString('\n')
 		if err != nil {log.Fatalf("receiving: %v",err)}
		if strings.TrimSpace(netData) == "STOP" {
			log.Println("Received STOP -- Exiting TCP server!")
			os.Exit(1)
		}

		fmt.Print("s-> ", netData)
		myTime := time.Now().Format(time.RFC3339) + "\n"
		con.Write([]byte(myTime))
	}
}

	func PrintCon(con *net.Conn) {

    loc := (*con).LocalAddr()
    fmt.Printf("Src Adr: %s %s\n", loc.Network(), loc.String())
    rem :=(*con).RemoteAddr()
    fmt.Printf("Dest Adr: %s %s\n", rem.Network(), rem.String())

}
