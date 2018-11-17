package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
			showLogs()
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
		Monitor(readFileReturnsArray("sites.txt"))
	case 2:
		CustomizedMonitor(config(), readFileReturnsArray("sites.txt"))
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
		var status bool
		response, err := http.Get(url)
		catError(err)

		if response.StatusCode == 200 {
			fmt.Println("Application on", url, "is UP and RUNING!!!!!")
			status = true
			writeFile("log.txt", url, status, err)
		} else {
			fmt.Println("ERROR on, ", url, "Something went Wrong!!!")
			status = false
			writeFile("log.txt", url, status, err)
		}
	}
}

func catError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func readFileReturnsArray(fileName string) []string {
	var urls []string

	file, err := os.Open(fileName)
	catError(err)

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		urls = append(urls, line)
	}
	file.Close()

	return urls
}

func readFile(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	catError(err)

	fmt.Println(string(file))
}

func writeFile(fileName string, url string, status bool, err error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	catError(err)

	file.WriteString("URL ==> " + url + ", active => " + strconv.FormatBool(status) + ", at => " + time.Now().Format("02/01/2006 15:04:05") + "\n")

	file.Close()
}

func showLogs() {
	fmt.Println("Showing Log...")
	readFile("log.txt")
}
