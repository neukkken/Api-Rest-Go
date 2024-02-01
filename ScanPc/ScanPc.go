package main

import (
	"fmt"
	"net"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func ScanPc(){
	
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error al obtener direcciones IP:", err)
		return
	}

	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println("Error al obtener información de la placa madre:", err)
		return
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error al obtener información de la memoria RAM:", err)
		return
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Println("Error al obtener información de la CPU:", err)
		return
	}

	for _, info := range cpuInfo {
		fmt.Println("------------------------------")
		fmt.Printf("CPU: %s\n", info.ModelName)
		fmt.Printf("Número de núcleos: %d\n", info.Cores)
		fmt.Println("Frecuencia: ", info.Mhz, "Mhz")
		fmt.Println("------------------------------")
	}

	fmt.Printf("Memoria RAM Total: %v GB\n", memInfo.Total/1024/1024/1024)
	fmt.Printf("Memoria RAM Libre: %v GB\n", memInfo.Free/1024/1024/1024)
	fmt.Printf("Memoria RAM Usada: %v GB\n", memInfo.Used/1024/1024/1024)
	fmt.Printf("Porcentaje de Uso: %f%%\n", memInfo.UsedPercent)
	fmt.Println("------------------------------")
	fmt.Println("Nombre del equipo:", hostInfo.Hostname)

	fmt.Println("------------------------------")
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("Dirección IP:", ipnet.IP.String())
			}
		}
	}
	fmt.Println("------------------------------")

}


func main() {

	ScanPc()

}
