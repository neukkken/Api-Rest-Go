package main

import (
	"fmt"
	"net"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func ScanTotalRam() uint64 {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error al obtener información de la memoria RAM:", err)
	}

	var totalRam = memInfo.Total / 1024 / 1024 / 1024

	return totalRam

}

func ScanFreeRam() uint64 {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error al obtener información de la memoria RAM:", err)
	}

	var freeRam = memInfo.Free / 1024 / 1024 / 1024

	return freeRam
}

func ScanUsedRam() uint64 {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error al obtener información de la memoria RAM:", err)
	}

	var usedRam = memInfo.Used / 1024 / 1024 / 1024

	return usedRam
}

func ScanPercentRam() float64 {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error al obtener información de la memoria RAM:", err)
	}

	var percentRam = memInfo.UsedPercent

	return percentRam

}

func ScanCpuName() string {
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Println("Error al obtener información de la CPU:", err)
	}

	var cpuName = cpuInfo[0].ModelName

	return cpuName

}

func ScanCpuCores() int32 {
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Println("Error al obtener información de la CPU:", err)
	}
	var cpuCores = cpuInfo[0].Cores / 2

	return cpuCores

}

func ScanCpuCache() int32 {
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Println("Error al obtener información de la CPU:", err)
	}
	var cpuCores = cpuInfo[0].CacheSize

	return cpuCores

}

func ScanCpuThreads() int32 {
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Println("Error al obtener información de la CPU:", err)
	}
	var cpuThreads = cpuInfo[0].Cores

	return cpuThreads

}

func ScanCpuFrecuency() float64 {
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Println("Error al obtener información de la CPU:", err)
	}
	var cpuFrecuency = cpuInfo[0].Mhz

	return cpuFrecuency

}

func ScanHostName() string {
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println("Error al obtener información de la placa madre:", err)
	}

	var hostName = hostInfo.Hostname

	return hostName
}

func ScanIps() string {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error al obtener direcciones IP:", err)
	}

	var IPS string

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {

				IPS = IPS + ipnet.IP.String() + "/"
			}
		}
	}

	return IPS
}

func ScanPc() {
	fmt.Println("------------------------------")
	ScanCpuName()
	fmt.Println("------------------------------")

	fmt.Println("------------------------------")
	ScanHostName()
	fmt.Println("------------------------------")
	ScanIps()
	fmt.Println("------------------------------")
}

func main() {
	fmt.Println(ScanCpuCache())

}
