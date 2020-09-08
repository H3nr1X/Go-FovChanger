package main

import (
    "fmt"
    "ReadWriteMemory"
    "log"
)
const (
    m_iFOV          = 0x31E4
    dwLocalPlayer   = 0xD3FC5C

)
func main() {
	process, err := ReadWriteMemory.ProcessByName("csgo")
	if err != nil {
		log.Panicf("csgo.exe not found. Error: %s", err.Error())
	}
    client := process.Modules["client.dll"].ModBaseAddr
    fmt.Print("Enter fov: ")
    var fov int
    fmt.Scanln(&fov)
	if err != nil {
		fmt.Println(err)
    }
    for {
    LocalBase, _ := process.ReadIntPtr(client + dwLocalPlayer)
    process.WriteInt(LocalBase + m_iFOV, fov)
    }
}