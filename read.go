package main

import (
	"fmt"
	"log"
	"os"
)

func ReadArg() string {
	arg := os.Args[1:]

	if len(arg) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(0)
	}

	port := "8989"

	if len(arg) == 1 {
		port = arg[0]
		for _, c := range port {

			if c > '9' || c < '0' {
				fmt.Println("[USAGE]: ./TCPChat $port")
				os.Exit(0)
			}
		}
	}
	return port
}

func Greating() string {
	dat, err := os.ReadFile("greating.txt")
    if err != nil {
		log.Println("Failed to load greating")
		return ""
	}
	return string(dat)
}

func Atoi(s string) int {
	res := 0
	m := 1
	for i := len(s) - 1; i >= 0; i-- {
		if rune(s[i])-'0' > 9 || rune(s[i])-'0' < 0 {
			return -1
		}
		res += m * int(rune(s[i])-'0')
		m *= 10
	}
	return res
}
