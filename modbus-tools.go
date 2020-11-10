package modbustools

import (
	"github.com/goburrow/modbus"
	"encoding/binary"
	"github.com/goburrow/serial"
	"time"
)

const (
	defaultSerialIdleTimeout = 60
	defaultAddress = "/dev/tty.usbserial-1410"
)
// Client from modbus
type Client modbus.Client

type serialConfig struct {
	// Serial port configuration.
	serial.Config
	IdleTimeout time.Duration
}

func setupConfig(address string, baudRate int, dataBits int, stopBits int, parity string, timeout int64) serialConfig{
	return serialConfig{
		Config: serial.Config{
			Address: address,
			BaudRate: baudRate,
			// Data bits: 5, 6, 7 or 8 (default 8)
			DataBits: dataBits,
			// Stop bits: 1 or 2 (default 1)
			StopBits: stopBits,
			// Parity: N - None, E - Even, O - Odd (default E)
			// (The use of no parity requires 2 stop bits.)
			Parity: parity,
			// Read (Write) timeout.
			Timeout: time.Duration(timeout) * time.Second,
		},
		IdleTimeout: time.Duration(defaultSerialIdleTimeout) * time.Second,
	}
}

func typicalConfig() serialConfig{
	return setupConfig(defaultAddress, 4800, 8, 1, "N", 3)
}

func setupHandler(config serialConfig) *modbus.RTUClientHandler{
	handler := modbus.NewRTUClientHandler(config.Address);
	handler.Timeout = config.Timeout
	handler.IdleTimeout = config.IdleTimeout
	return handler;
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func convertParity(parity int) string{
	if parity == 0 {
		return "N"
	}
	if parity == 1 {
		return "O"
	}
	if parity == 2 {
		return "E"
	}
	return "N"
}

func uintsToBytes(vs []uint32) []byte {
	buf := make([]byte, len(vs)*4)
	for i, v := range vs {
		binary.BigEndian.PutUint32(buf[i*4:], v)
	}
	return buf
}

func uints16ToBytes(vs []uint16) []byte {
	buf := make([]byte, len(vs)*2)
	for i, v := range vs {
		binary.BigEndian.PutUint16(buf[i*2:], v)
	}
	return buf
}
