package autofunc

import (
	"fmt"
	"net"
	"runtime"
	"syscall"
	"time"
)

func FaisDodo() {
	time.Sleep(time.Second * 5)
}

func ProgMain(pid int) {

	startCpuTotalTime, _ := TempTotalCpu()
	startProcessTime, _ := TempProcessCPU(pid)
	startCalculMemory := MemoryUse()

	var data []byte
	for i := 0; i < 1000000; i++ {
		data = append(data, make([]byte, 1000)...)
		_ = i * i
		if len(data) > 1024*1024*10 {
			break
		}

	}

	endCpuTotalTime, _ := TempTotalCpu()
	endProcessTime, _ := TempProcessCPU(pid)
	endCalculMemory := MemoryUse()

	finalCpuTime := endCpuTotalTime - startCpuTotalTime
	finalProcessTime := endProcessTime - startProcessTime
	finalMemoryCalcul := endCalculMemory - startCalculMemory

	cpuUsagePerCore := (float64(finalProcessTime) / float64(finalCpuTime)) * 100

	fmt.Println("Temps total CPU :\t\t\t", finalCpuTime)
	fmt.Println("Temps CPU in process  :\t\t\t", finalProcessTime)
	fmt.Printf("Usage of CPU :\t\t\t\t %.2f%%\n", cpuUsagePerCore)
	fmt.Printf("consommation memory of process :\t %f Mo\n\n", float64(finalMemoryCalcul)/(1024*1024))

}

func TempTotalCpu() (time.Duration, error) {
	var usage syscall.Rusage
	if err := syscall.Getrusage(syscall.RUSAGE_SELF, &usage); err != nil {
		return 0, err
	}
	return time.Duration(usage.Utime.Nano() + usage.Stime.Nano()), nil
}

func TempProcessCPU(pid int) (time.Duration, error) {
	var usage syscall.Rusage
	if err := syscall.Getrusage(syscall.RUSAGE_SELF, &usage); err != nil {
		return 0, err
	}
	return time.Duration(usage.Utime.Nano() + usage.Stime.Nano()), nil
}

func MemoryUse() uint64 {
	var memory runtime.MemStats
	runtime.ReadMemStats(&memory)
	return memory.Alloc
}

func InfoNetwork() {
	// interface := net.Interfaces()
	getInterface, err := net.Interfaces()
	if err != nil {
		fmt.Println("Erreur de récupération des interfaces réseaux ! .", err)
		return
	}

	for _, interff := range getInterface {

		fmt.Println("Name of interfaces : ", interff.Name)
		// fmt.Println("Name of Adress MAC : ", interff.HardwareAddr)

		addrs, _ := interff.Addrs()
		for _, addr := range addrs {
			fmt.Println("Adresse IP : ", addr)
		}
	}
}

func InfoMacAdress(interfaceName string) (string, error) {
	itface, err := net.InterfaceByName((interfaceName))
	if err != nil {
		return "", err
	}

	macAdresse := itface.HardwareAddr.String()
	return macAdresse, nil
}
