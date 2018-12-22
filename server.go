package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	// arguments := os.Args
	// if len(arguments) == 1 {
	// 	fmt.Println("Please provide a port number!")
	// 	return
	// }

	PORT := ":3001" //madhuxx + arguments[1]
	fmt.Println("Starting the Serer on PORT", PORT)
	fmt.Println("Before Listen")
	listner, err := net.Listen("tcp", PORT)
	fmt.Println("After Listen")
	if err != nil {
		fmt.Println("Error While Listne:", err)
		return
	}
	fmt.Println("Before defer")
	//We don't need this as we don't want to kill the server if any of the
	//client goes offline. we want to run the sever for other clients.
	//defer listner.Close()
	//rand.Seed(time.Now().Unix())
	fmt.Println("Before For")
	for {
		fmt.Println("Before Accept")
		conn, err := listner.Accept()
		fmt.Println("After Accept")
		if err != nil {
			fmt.Println("Error While Accept:", err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	for {
		fmt.Println("Inside for of handleConnection :- before newReader for:", conn.RemoteAddr().String())
		netData, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Inside for of handleConnection :- After newReader")
		if err != nil {
			fmt.Println("Error While Read:", err)
			return
		}
		fmt.Println("Inside for of handleConnection :- Before Trim")
		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		//generate random number
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		randInt := r.Intn(100000)

		result := "Response from Server is :" + temp + ":with Random #: " + strconv.Itoa(randInt) + "\n"
		fmt.Println("Inside for of handleConnection :- Before Write")
		fmt.Println("Going to send the response to clien:", result)
		conn.Write([]byte(string(result)))
		fmt.Println("Inside for of handleConnection :- After Write")
	}
	conn.Close()
}
