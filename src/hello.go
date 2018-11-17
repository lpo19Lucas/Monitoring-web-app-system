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
	// var frutas [4]string
	// frutas[0] = "Abacaxi"
	// frutas[1] = "Laranja"
	// frutas[2] = "Morango"
	// fmt.Println(time.Duration(3))
	// fmt.Println0(readFile("sites.txt"))
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
	// sites := []string{"http://serene-meadow-32620.herokuapp.com/api/notes", "http://www.catho.com.br", "http://www.alura.com.br"}

	fmt.Println("Before begin chose each type of monitoring must be executed!")
	fmt.Println("1- Default Monitoring(Only one time each application)")
	fmt.Println("2- Configure Monitoring")

	comand := catComand()

	switch comand {
	case 1:
		Monitor(readFile("sites.txt"))
	case 2:
		CustomizedMonitor(configMonitoring(), readFile("sites.txt"))
	default:
		fmt.Println(comand, "It's NOT a valid option!")
		os.Exit(-1)
	}

}

func configMonitoring() int {

	fmt.Println("Enter the time interval you want to monitor as apps")
	times := catComand()
	return times

}

func CustomizedMonitor(interval int, sites []string) {

	fmt.Println("Starting Monitoring...")
	for {
		for _, site := range sites {
			response, err := http.Get(site)

			catError(err)

			if response.StatusCode == 200 {
				fmt.Println("Application on", site, "is UP and RUNING!!!!!")
			} else {
				fmt.Println("ERROR on, ", site, "Something went Wrong!!!")
			}
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func Monitor(sites []string) {
	fmt.Println("Starting Monitoring...")
	for _, site := range sites {
		response, err := http.Get(site)

		catError(err)

		if response.StatusCode == 200 {
			fmt.Println("Application on", site, "is UP and RUNING!!!!!")
		} else {
			fmt.Println("ERROR on, ", site, "Something went Wrong!!!")
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
	var sites []string

	response, err := os.Open(file)
	catError(err)

	reader := bufio.NewReader(response)

	for {
		linha, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)
	}

	response.Close()

	return sites
}
