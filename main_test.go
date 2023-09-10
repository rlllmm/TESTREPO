package main

import (
	"autofunc/autofunc"
	"os"
	"testing"
	"time"
)

func TestMemoryUse(t *testing.T) {
	memory := autofunc.MemoryUse()
	if memory <= 0 {
		t.Errorf("Expected positive memory usage, got %d", memory)
	}
}

func TestTempTotalCpu(t *testing.T) {
	startTime, _ := autofunc.TempTotalCpu()
	// Simulate some CPU-bound work.
	time.Sleep(100 * time.Millisecond)
	endTime, _ := autofunc.TempTotalCpu()

	if endTime <= startTime {
		t.Errorf("Expected endTime to be greater than startTime, got endTime=%d, startTime=%d", endTime, startTime)
	}
}

func TestTempProcessCPU(t *testing.T) {
	pid := os.Getpid()
	startTime, _ := autofunc.TempProcessCPU(pid)
	// Simulate some CPU-bound work.
	time.Sleep(100 * time.Millisecond)
	endTime, _ := autofunc.TempProcessCPU(pid)

	if endTime <= startTime {
		t.Errorf("Expected endTime to be greater than startTime, got endTime=%d, startTime=%d", endTime, startTime)
	}
}
