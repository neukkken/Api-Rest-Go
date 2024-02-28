package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"bytes"
	
	"github.com/gorilla/mux"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/elastic/go-elasticsearch/v8"
    
	
	
	"crypto/tls"
)

type pc struct {
	ID           int     `json:ID`
	PcName       string  `json:PcName`
	CpuName      string  `json:CpuName`
	CpuCores     int32   `json:CpuCores`
	CpuThreads   int32   `json:CpuThreads`
	CpuFrecuency float64 `json:CpuFrecuency`
	PcIPS        string  `json:PcIPS`
	TotalRam     uint64  `json:TotalRam`
	FreeRam      uint64  `json:FreeRam`
	UsedRam      uint64  `json:UsedRam`
	PercentRam   float64 `json:PercentRam`
}

type allPcs []pc

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

				IPS = IPS + "/" + ipnet.IP.String() + "   "
			}
		}
	}
	return IPS
}

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

var pcs = allPcs{
	{
		ID:           1,
		PcName:       ScanHostName(),
		CpuName:      ScanCpuName(),
		CpuCores:     ScanCpuCores(),
		CpuThreads:   ScanCpuThreads(),
		CpuFrecuency: ScanCpuFrecuency(),
		PcIPS:        ScanIps(),
		TotalRam:     ScanTotalRam(),
		FreeRam:      ScanFreeRam(),
		UsedRam:      ScanUsedRam(),
		PercentRam:   ScanPercentRam(),
	},
}

func getPcs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pcs)
	w.Header().Set("Content-Type", "application/json")
}

func createPc(w http.ResponseWriter, r *http.Request) {
	var newPc pc

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid task")
	}

	json.Unmarshal(reqBody, &newPc)

	newPc.ID = len(pcs) + 1
	pcs = append(pcs, newPc)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPc)

}

func getPc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pcID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, pc := range pcs {
		if pc.ID == pcID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pc)
		}
	}

}

func deletePc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pcID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "invalid ID")
		return
	}

	for i, pc := range pcs {
		if pc.ID == pcID {
			pcs = append(pcs[:i], pcs[i+1:]...)
			fmt.Fprintf(w, "The pcs was delete succesfully")
		}
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to api")
}

func main() {

	// router := mux.NewRouter().StrictSlash(true)

	// router.HandleFunc("/", indexRoute)
	// router.HandleFunc("/pcs", getPcs).Methods("GET")
	// router.HandleFunc("/pcs", createPc).Methods("POST")
	// router.HandleFunc("/pcs/{id}", getPc).Methods("GET")
	// router.HandleFunc("/pcs/{id}", deletePc).Methods("DELETE")
	// log.Fatal(http.ListenAndServe(":3000", router))                                                
// Aquí está tu configuración inicial para la conexión a Elasticsearch
cfg := elasticsearch.Config{
	Addresses: []string{"https://localhost:9200"},
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,    
		},
	},
	Username: "elastic",
	Password: "xIkLoSDW8Z5EcuXpsXIv",
}

// Crear un cliente Elasticsearch
es, err := elasticsearch.NewClient(cfg)
if err != nil {
	log.Fatalf("Error creating the client: %s", err)
}

// Datos para indexar en Elasticsearch
data := map[string]interface{}{
	"cpu": ScanCpuName(),
	"ram": ScanTotalRam(),
	"ips": ScanIps(),
}

// Convertir los datos a formato JSON
jsonData, err := json.Marshal(data)
if err != nil {
	log.Fatalf("Error al convertir los datos a JSON: %s", err)
}

// Indexar el documento en Elasticsearch
res, err := es.Index("pcs", bytes.NewReader(jsonData))
if err != nil {
	log.Fatalf("Error indexing document: %s", err)
}
defer res.Body.Close()

// Verificar el código de estado de la respuesta
if res.IsError() {
	log.Fatalf("Error al indexar documento: %s", res.Status()) 
}

log.Println("Documento indexado exitosamente!")


}
