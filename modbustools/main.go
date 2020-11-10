package main

import (
	"encoding/binary"
	"fmt"

	modbustools "github.com/carlosfrutos/modbus-tools"
	"github.com/goburrow/modbus"
)


func main(){
	config := modbustools.SetupConfig("/dev/tty.usbserial-1410", 4800, 8, 2, "N", 10)
	handler := modbustools.SetupHandler(config)

	handler.SlaveId = 1
	err := handler.Connect()
	defer handler.Close()
	modbustools.Check(err)

	client := modbus.NewClient(handler)
	results, err := client.ReadHoldingRegisters(uint16(0), 1)
	modbustools.Check(err)
	if len(results) > 0 {
		data := binary.BigEndian.Uint16(results)
		fmt.Printf("Register data: %v", data)
	} else {
		fmt.Printf("Register data empty: %v", results)
	}
    
}