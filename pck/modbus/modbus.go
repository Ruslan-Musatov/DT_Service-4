package modbus

import (
    "log"

    "github.com/goburrow/modbus"
)

// ReadTemperature читает температуру с устройства Modbus и возвращает результаты в виде среза float32.
func ReadTemperature(port string, baudRate int, slaveID byte, startAddress, endAddress uint16) ([]float32, error) {
    handler := modbus.NewRTUClientHandler(port)
    handler.BaudRate = baudRate
    handler.DataBits = 8
    handler.Parity = "N"
    handler.StopBits = 1
    handler.SlaveId = slaveID
    handler.Timeout = 1

    err := handler.Connect()
    if err != nil {
        return nil, err
    }
    defer handler.Close()

    client := modbus.NewClient(handler)

    var temperatures []float32
    for address := startAddress; address <= endAddress; address++ {
        results, err := client.ReadInputRegisters(address, 1)
        if err != nil {
            log.Printf("Ошибка при чтении с адреса %d: %v", address, err)
            continue
        }

        if len(results) >= 2 {
            temperature := float32(uint16(results[0])<<8|uint16(results[1])) / 100.0
            temperatures = append(temperatures, temperature)
        }
    }

    return temperatures, nil
}
