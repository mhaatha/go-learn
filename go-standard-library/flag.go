package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "root", "database username")
	password := flag.String("password", "root", "database password")
	host := flag.String("host", "root", "database host")
	port := flag.Int("port", 0, "database port")
	isLogin := flag.Bool("login", false, "is user login?")

	flag.Parse()

	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Host:", *host)
	fmt.Println("Port:", *port)
	fmt.Println("isLogin:", *isLogin)
}
