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
		comand := catchComand()
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

func catchComand() int {
	var selected int
	fmt.Scan(&selected)

	return selected
}

func initializeMonitoring() {
	fmt.Println("Before begin chose each type of monitoring must be executed!")
	fmt.Println("1- Default Monitoring(Only one time each application)")
	fmt.Println("2- Configure Monitoring")

	comand := catchComand()

	switch comand {
	case 1:
		Monitor(readFileReturnsArray("sites.txt"))
	case 2:
		CustomizedMonitor(conf1(), conf2(), readFileReturnsArray("sites.txt"))
	default:
		fmt.Println(comand, "It's NOT a valid option!")
		os.Exit(-1)
	}
}

func conf1() int {
	fmt.Println("How many times do you want to monitor as apps")
	monitoringTimes := catchComand()

	return monitoringTimes
}

func conf2() int {
	fmt.Println("Enter the time-out value(seconds) you want to monitor your apps")
	timeOut := catchComand()

	return timeOut
}

func CustomizedMonitor(monitorTimes int, timeOut int, urls []string) {
	fmt.Println("Starting Monitoring...")

	for i := 0; i < monitorTimes; i++ {
		for _, url := range urls {
			response, err := http.Get(url)
			catchError(err)

			if response.StatusCode == 200 {
				fmt.Println("Application on", url, "is UP and RUNING!!!!!")
				writeLog(url, response.StatusCode, err)

			} else {
				fmt.Println("ERROR on, ", url, "Something went Wrong!!!")
				writeLog(url, response.StatusCode, err)

			}
		}
		time.Sleep(time.Duration(timeOut) * time.Second)
	}
}

func Monitor(urls []string) {
	fmt.Println("Starting Monitoring...")

	for _, url := range urls {
		response, err := http.Get(url)
		catchError(err)

		if response.StatusCode == 200 {
			fmt.Println("Application on", url, "is UP and RUNING!!!!!")
			writeLog(url, response.StatusCode, err)

		} else {
			fmt.Println("ERROR on, ", url, "Something went Wrong!!!")
			writeLog(url, response.StatusCode, err)

		}
	}
}

func catchError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func readFileReturnsArray(fileName string) []string {
	var urls []string

	file, err := os.Open(fileName)
	catchError(err)

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
	catchError(err)

	fmt.Println(string(file))
}

func writeLog(url string, statusCode int, err error) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	catchError(err)

	file.WriteString("URL ==> " + url + ", statusCode => " + strconv.FormatInt(int64(statusCode), 10) + ", at => " + time.Now().Format("02/01/2006 15:04:05") + "\n")

	file.Close()
}

func showLogs() {
	fmt.Println("Showing Log...")
	readFile("log.txt")
}
