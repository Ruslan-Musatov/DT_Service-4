package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io"
    "os"
    
)

type YourStruct struct {
    // Определите поля структуры, которые соответствуют вашему формату данных
    Field1 string `json:"field1"`
    Field2 int    `json:"field2"`
   
}

func main() {
    // Открываем файл RTA
    rtaFile, err := os.Open("C:/Users/musatov/Desktop/example_0.rta")
    if err != nil {
        fmt.Println("Ошибка при открытии файла RTA:", err)
        return
    }
    defer rtaFile.Close()

    // Создаем файл TXT
    txtFile, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Ошибка при создании файла TXT:", err)
        return
    }
    defer txtFile.Close()

    // Создаем буферизированные ридеры и райтеры
    reader := bufio.NewReader(rtaFile)
    writer := bufio.NewWriter(txtFile)

    // Читаем и записываем содержимое файла
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println("Ошибка при чтении файла RTA:", err)
            return
        }

        _, err = writer.WriteString(line)
        if err != nil {
            fmt.Println("Ошибка при записи в файл TXT:", err)
            return
        }
    }

    // Сбрасываем буфер на диск
    writer.Flush()

    // Прочитать файл output.txt и преобразовать его в JSON
    jsonFile, err := os.Open("output.txt")
    if err != nil {
        fmt.Println("Ошибка при открытии файла output.txt:", err)
        return
    }
    defer jsonFile.Close()

    var jsonData []YourStruct 
    scanner := bufio.NewScanner(jsonFile)
    for scanner.Scan() {
        // Предполагается, что каждая строка файла содержит JSON объект
        line := scanner.Text()
        var data YourStruct
        err := json.Unmarshal([]byte(line), &data)
        if err != nil {
            fmt.Println("Ошибка при разборе JSON:", err)
            continue
        }
        jsonData = append(jsonData, data)
    }

    // Проверка ошибок сканера
    if err := scanner.Err(); err != nil {
        fmt.Println("Ошибка сканирования файла:", err)
        return
    }

    // Вывести полученный JSON
    jsonOutput, err := json.Marshal(jsonData)
    if err != nil {
        fmt.Println("Ошибка при преобразовании в JSON:", err)
        return
    }

    fmt.Println(string(jsonOutput))
    fmt.Println("Конвертация завершена успешно!")
}
