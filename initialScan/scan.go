package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"runtime"

	"nmap"
)

func main() {
	//Asks user for box name
	fmt.Printf("Enter box name: ")
	var boxName string
	fmt.Scanln(&boxName)

	//Asks user for ip address
	fmt.Printf("Enter ip address: ")
	var ipAddress string
	fmt.Scanln(&ipAddress)

	choices(boxName, ipAddress)
}

func checkError(e error) {
	//Performs error checking
	if e != nil {
		log.Fatal(e)
	}
}
func choices(boxName, ipAddress string) {
	fmt.Printf(`
 /$$   /$$ /$$$$$$$$ /$$$$$$$
| $$  | $$|__  $$__/| $$__  $$
| $$  | $$   | $$   | $$  \ $$  /$$$$$$$  /$$$$$$$  /$$$$$$  /$$$$$$$
| $$$$$$$$   | $$   | $$$$$$$  /$$_____/ /$$_____/ |____  $$| $$__  $$
| $$__  $$   | $$   | $$__  $$|  $$$$$$ | $$        /$$$$$$$| $$  \ $$
| $$  | $$   | $$   | $$  \ $$ \____  $$| $$       /$$__  $$| $$  | $$
| $$  | $$   | $$   | $$$$$$$/ /$$$$$$$/|  $$$$$$$|  $$$$$$$| $$  | $$
|__/  |__/   |__/   |_______/ |_______/  \_______/ \_______/|__/  |__/` + "\n")

	fmt.Printf("\t\tby 3nt3r\n")
	fmt.Printf("\n\t{1} Nmap scans\n")
	fmt.Printf("\t{2} Gobuster\n")
	fmt.Printf("\n>>>")
	var choice string
	fmt.Scanln(&choice)

	if choice == "1" {
		nmapScans(boxName, ipAddress)
	} else if choice == "2" {
		gobusterScans()
	}

}

func nmapScans(boxName, ipAddress string) {


	fmt.Printf("Please choose a scan: ")
	fmt.Printf("\n\t{1} Basic Scan ")
	fmt.Printf("\n\t{2} Medium Scan ")
	fmt.Printf("\n\t{3} Detailed Scan")
	fmt.Printf("\n\t{4} Basic Scan & Detailed Scan")
	fmt.Printf("\n\t{5} Default Vulnerability Script")
	fmt.Printf("\n\t{0} Return to Menu")

	fmt.Printf("\n>>>")
	var choice string
	fmt.Scanln(&choice)

	if choice == "1" {
		nmap.BasicNmapScan(boxName, ipAddress)
	} else if choice == "2" {
		nmap.MediumNmapScan(boxName, ipAddress)
	} else if choice == "3" {
		nmap.FullNmapScan(boxName, ipAddress)
	} else if choice == "4" {
		runtime.GOMAXPROCS(2)

		var wg sync.WaitGroup
		wg.Add(2)
		//fmt.Printf("Starting Basic Scan...\n")
		go func() {
			defer wg.Done()
			nmap.BasicNmapScan(boxName, ipAddress)
		//	fmt.Printf("Finished Basic Scan.")
		}()

		//fmt.Printf("Starting Detailed Scan...")
		go func() {
			defer wg.Done()
			nmap.FullNmapScan(boxName, ipAddress)
			//fmt.Printf("Finished Detailed Scan.")
		}()

		wg.Wait()
	} else if choice == "5" {
		nmap.DefaultVulnScript(boxName, ipAddress)
	} else if choice == "0" {
		choices(boxName, ipAddress)
	}
}

func gobusterScans() {

}

/*func info() {
	//Gets the current working directory
	dir, err := os.Getwd()
	checkError(err)

	fmt.Printf("All output will be saved to " + dir + "\n")
	fmt.Printf("Save output to a different directory? y/n: ")
	var saveLocation string
	fmt.Scanln(&saveLocation)
	//Takes the user response and performs
	//Necessary action
	if saveLocation == "y"{
		fmt.Printf("Enter the new save location: ")
		var newLocation string
		fmt.Scanln(&newLocation)
		checkDir(newLocation)
	} else if saveLocation == "n" {
		input()
	} else {
		fmt.Printf("Not a valid option\n")
	}
}*/

func checkDir(directory string) {
	//Checks if directory path is valid
	src, err := os.Stat(directory)
	checkError(err)
	//checks if path entered is really a directory
	if !src.IsDir() {
		fmt.Printf("Source is not a directory.")
		os.Exit(1)
	} else if src.IsDir() {
		//Changes woring directory
		os.Chdir(directory)
	}
}
