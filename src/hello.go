package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	Monitorising()
}

func Monitorising() {
	for {
		showMenu()
		comand := catComand()
		switch comand {
		case 1:
			initializeMonitoring()
		case 2:
			fmt.Println("Showing Log...")
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println(comand, "It's NOT a valid option!")
			os.Exit(-1)
		}
	}
}

func showMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func catComand() int {
	var selected int
	fmt.Scan(&selected)

	return selected
}

func initializeMonitoring() {
	fmt.Println("Before begin chose each type of monitoring must be executed!")
	fmt.Println("1- Default Monitoring(Only one time each application)")
	fmt.Println("2- Configure Monitoring")

	comand := catComand()

	switch comand {
	case 1:
		Monitor(readFile("sites.txt"))
	case 2:
		CustomizedMonitor(config(), readFile("sites.txt"))
	default:
		fmt.Println(comand, "It's NOT a valid option!")
		os.Exit(-1)
	}
}

func config() int {
	fmt.Println("Enter the time-out value(seconds) you want to monitor as apps")
	time := catComand()

	return time
}

func CustomizedMonitor(timeOut int, urls []string) {
	fmt.Println("Starting Monitoring...")

	for {
		for _, url := range urls {
			response, err := http.Get(url)
			catError(err)

			if response.StatusCode == 200 {
				fmt.Println("Application on", url, "is UP and RUNING!!!!!")
			} else {
				fmt.Println("ERROR on, ", url, "Something went Wrong!!!")
			}
		}
		time.Sleep(time.Duration(timeOut) * time.Second)
	}
}

func Monitor(urls []string) {
	fmt.Println("Starting Monitoring...")

	for _, url := range urls {
		response, err := http.Get(url)
		catError(err)

		if response.StatusCode == 200 {
			fmt.Println("Application on", url, "is UP and RUNING!!!!!")
		} else {
			fmt.Println("ERROR on, ", url, "Something went Wrong!!!")
		}
	}
}

func catError(err error) {
	if err != nil {
		fmt.Println(err)
	}

	return
}

func readFile(file string) []string {
	var urls []string

	response, err := os.Open(file)
	catError(err)

	reader := bufio.NewReader(response)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		urls = append(urls, line)
	}
	response.Close()

	return urls
}
