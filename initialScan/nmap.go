package nmap

import (
  "context"
  "log"
  "fmt"
  "time"

  "github.com/Ullaakut/nmap"
)

func checkError(e error) {
	//Performs error checking
	if e != nil {
		log.Fatal(e)
	}
}
func BasicNmapScan(boxName, ipAddress string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	//Equivalent to /usr/local/bin/nmap -sV -sC -T4 ipAddress
	//With a 5 minute timeout
	scanner, err := nmap.NewScanner(
		nmap.WithTimingTemplate(4),
		nmap.WithTargets(ipAddress),
		nmap.WithContext(ctx),
	)
	checkError(err)

	result, err := scanner.Run()
	checkError(err)

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name, port.Service.Version, port.Service.ExtraInfo)
		}
	}
	fmt.Printf("Nmap done: %d hosts up scanned in %3f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)

}

func MediumNmapScan(boxName, ipAddress string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	//Equivalent to /usr/local/bin/nmap -sV -sC -T4 ipAddress
	//With a 5 minute timeout
	scanner, err := nmap.NewScanner(
		nmap.WithServiceInfo(),
		nmap.WithDefaultScript(),
		nmap.WithTimingTemplate(4),
		nmap.WithTargets(ipAddress),
		nmap.WithContext(ctx),
	)
	checkError(err)

	result, err := scanner.Run()
	checkError(err)

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name, port.Service.Version, port.Service.ExtraInfo)
		}
	}
	fmt.Printf("Nmap done: %d hosts up scanned in %3f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)

}

func FullNmapScan(boxName, ipAddress string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	//Equivalent to /usr/local/bin/nmap -sV -sC -T4 ipAddress
	//With a 5 minute timeout
	scanner, err := nmap.NewScanner(
		nmap.WithServiceInfo(),
		nmap.WithDefaultScript(),
		nmap.WithTimingTemplate(4),
		nmap.WithPorts("-"),
		nmap.WithTargets(ipAddress),
		nmap.WithContext(ctx),
	)
	checkError(err)

	result, err := scanner.Run()
	checkError(err)

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name, port.Service.Version, port.Service.ExtraInfo)
		}
	}
	fmt.Printf("Nmap done: %d hosts up scanned in %3f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)

}
func DefaultVulnScript(boxName, ipAddress string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	//Equivalent to /usr/local/bin/nmap -sV -sC -T4 ipAddress
	//With a 5 minute timeout
	scanner, err := nmap.NewScanner(
		nmap.WithScripts("vuln"),
		nmap.WithTargets(ipAddress),
		nmap.WithContext(ctx),
	)
	checkError(err)

	result, err := scanner.Run()
	checkError(err)

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name, port.Service.Version, port.Service.ExtraInfo)
		}
	}
	fmt.Printf("Nmap done: %d hosts up scanned in %3f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)

}
