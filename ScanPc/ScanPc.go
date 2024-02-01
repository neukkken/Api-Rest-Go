package main

import (
	"fmt"
	"net"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func ScanRam() {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error al obtener información de la memoria RAM:", err)
		return
	}

	fmt.Printf("Memoria RAM Total: %v GB\n", memInfo.Total/1024/1024/1024)
	fmt.Printf("Memoria RAM Libre: %v GB\n", memInfo.Free/1024/1024/1024)
	fmt.Printf("Memoria RAM Usada: %v GB\n", memInfo.Used/1024/1024/1024)
	fmt.Printf("Porcentaje de Uso: %f%%\n", memInfo.UsedPercent)
}

func ScanCpu() {
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Println("Error al obtener información de la CPU:", err)
		return
	}

	fmt.Println("------------------------------")
	fmt.Printf("CPU: %s\n", cpuInfo[0].ModelName)
	fmt.Printf("Número de núcleos: %d\n", cpuInfo[0].Cores)
	fmt.Println("Frecuencia: ", cpuInfo[0].Mhz, "Mhz")
	fmt.Println("------------------------------")
}

func ScanHostName() {
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println("Error al obtener información de la placa madre:", err)
		return
	}

	fmt.Println("------------------------------")
	fmt.Println("Nombre del equipo:", hostInfo.Hostname)

	fmt.Println("------------------------------")
}

func ScanIp() {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error al obtener direcciones IP:", err)
		return
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("Dirección IP:", ipnet.IP.String())
			}
		}
	}

}

func ScanPc() {

	ScanCpu()
	ScanRam()
	ScanHostName()
	ScanIp()

}

func main() {
	ScanPc()

}
