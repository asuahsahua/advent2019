package intcode

import (
	"sync"
)

type MachineState struct {
	StatusCode MachineStatusCode
	Lock *sync.RWMutex
}

func NewMachineState() *MachineState {
	return &MachineState{
		StatusCode: Suspended,
		Lock: &sync.RWMutex{},
	}
}

type MachineStatusCode int
const (
	Suspended MachineStatusCode = 0
	Running MachineStatusCode = 1
	OnFire MachineStatusCode = 2
)

func (ms *MachineState) Set(msc MachineStatusCode) {
	ms.Lock.Lock()
	ms.StatusCode = msc
	ms.Lock.Unlock()
}

func (ms *MachineState) Get() (msc MachineStatusCode) {
	ms.Lock.RLocker().Lock()
	msc = ms.StatusCode
	ms.Lock.RLocker().Unlock()
	return
}