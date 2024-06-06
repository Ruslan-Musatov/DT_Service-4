// Ваш файл main.go

package main

import (
	"fmt"

	"github.com/yourusername/modbusreader" // Замените "yourusername" на ваше имя пользователя GitHub или путь к пакету.
)

func main() {
	port := "COM5"
	baudRate := uint(115200)
	slaveID := uint8(17)
	startAddress := uint16(7)
	endAddress := uint16(12)

	temperatures, err := modbusreader.ReadTemperature(port, baudRate, slaveID, startAddress, endAddress)
	if err != nil {
		fmt.Printf("Ошибка при чтении температуры: %v\n", err)
		return
	}

	for i, temperature := range temperatures {
		fmt.Printf("Температура %d: %.2f\n", i+1, temperature)
	}
}
