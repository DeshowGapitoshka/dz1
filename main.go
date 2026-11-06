package main

import (
	driver "dz1/driver"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 5 {
		log.Println("Неверный набор аргументов")
		os.Exit(1)
	}
	if args[1] == "--impl" {
		if args[2] == "bst" {
			in, err := os.Open(args[3])
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			out, err := os.OpenFile(args[4], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			defer in.Close()
			driver.DriverBst(in, out, 1)
		}
		if args[2] == "hash" {
			in, err := os.Open(args[3])
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			out, err := os.OpenFile(args[4], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			defer in.Close()
			driver.DriverHash(in, out, 1)
		}
		if args[2] == "tree" {
			in, err := os.Open(args[3])
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			out, err := os.OpenFile(args[4], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			defer in.Close()
			driver.DriverRbt(in, out, 1)
		}
	} else if args[1] == "--benchmark" {
		if args[2] != "" && args[3] != "" {
			in, err := os.Open(args[3])
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			out, err := os.OpenFile(args[4], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			defer in.Close()
			driver.Benchmark(in, out)
		}
	} else {
		fmt.Println("Неверный набор аргументов")
	}

}
