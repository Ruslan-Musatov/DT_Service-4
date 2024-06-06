// Ваш файл modbus.go

package modbusreader

import (
	"log"

	"github.com/goburrow/modbus"
)

// ReadTemperature читает температуру с Modbus устройства в заданном диапазоне адресов.
func ReadTemperature(port string, baudRate uint, slaveID uint8, startAddress, endAddress uint16) ([]float32, error) {
	handler := modbus.NewRTUClientHandler(port)
	handler.BaudRate = int(baudRate) // Преобразуем uint в int
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
