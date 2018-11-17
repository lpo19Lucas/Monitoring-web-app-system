package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// var frutas [4]string
	// frutas[0] = "Abacaxi"
	// frutas[1] = "Laranja"
	// frutas[2] = "Morango"
	// fmt.Println(time.Duration(3))
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
			fmt.Println("Exibindo Log...")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println(comand, "Não é uma opção valida!")
			os.Exit(-1)
		}
	}
}

func showMenu() {
	fmt.Println("1- Inicar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")
}

func catComand() int {
	var selected int
	fmt.Scan(&selected)

	return selected
}

func initializeMonitoring() {
	sites := []string{"http://serene-meadow-32620.herokuapp.com/api/notes", "http://www.catho.com.br", "http://www.alura.com.br"}

	fmt.Println("Before begin chose each type of monitoring must be executed!")
	fmt.Println("1- Default Monitoring(Only one time each application)")
	fmt.Println("2- Configure Monitoring")

	comand := catComand()

	switch comand {
	case 1:
		Monitor(sites)
	case 2:
		CustomizedMonitor(configMonitoring(), sites)
	default:
		fmt.Println(comand, "Não é uma opção valida!")
		os.Exit(-1)
	}

}

func configMonitoring() int {

	fmt.Println("Digite de quanto em quanto tempo(segundos) você deseja monitorar as aplicações")
	times := catComand()
	return times

}

func CustomizedMonitor(interval int, sites []string) {

	fmt.Println("Iniciando Monitoramento...")
	for {
		for _, site := range sites {
			response, _ := http.Get(site)

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
	fmt.Println("Iniciando Monitoramento...")
	for _, site := range sites {
		response, _ := http.Get(site)

		if response.StatusCode == 200 {
			fmt.Println("Application on", site, "is UP and RUNING!!!!!")
		} else {
			fmt.Println("ERROR on, ", site, "Something went Wrong!!!")
		}
	}
}
