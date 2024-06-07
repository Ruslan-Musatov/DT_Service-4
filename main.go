package main
import (
    "encoding/json"
    "fmt"
    "net/http"
    //"time"
)
type Data struct {
    Address uint16  `json:"address"`
    Value   float64 `json:"value"`
}
func fetchDataFromServer() {
    for {
        data1, err := requestDataArray("http://172.27.69.227:5555/TempChanel")
        if err != nil {
            fmt.Println("Ошибка при получении данных из точки входа 1:", err)
        } else {
            fmt.Println("Полученные данные из точки входа 1:", data1)
        }
        data2, err := requestDataObject("http://172.27.69.227:5555/TempRadiator")
        if err != nil {
            fmt.Println("Ошибка при получении данных из точки входа 2:", err)
        } else {
            fmt.Println("Полученные данные из точки входа 2:", data2)
        }
        data3, err := requestDataObject("http://172.27.69.227:5555/TempCap")
        if err != nil {
            fmt.Println("Ошибка при получении данных из точки входа 3:", err)
        } else {
            fmt.Println("Полученные данные из точки входа 3:", data3)
        }
        // Добавьте больше вызовов requestData для других точек входа по вашему усмотрению
        //time.Sleep(1 * time.Second) // Ждем 5 секунд перед следующим запросом
    }
}
func requestDataArray(endpoint string) ([]Data, error) {
    resp, err := http.Get(endpoint)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("ошибка при получении данных: %s", resp.Status)
    }
    var data []Data
    err = json.NewDecoder(resp.Body).Decode(&data)
    if err != nil {
        return nil, fmt.Errorf("ошибка при разборе данных: %v", err)
    }
    return data, nil
}
func requestDataObject(endpoint string) (Data, error) {
    resp, err := http.Get(endpoint)
    if err != nil {
        return Data{}, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return Data{}, fmt.Errorf("ошибка при получении данных: %s", resp.Status)
    }
    var data Data
    err = json.NewDecoder(resp.Body).Decode(&data)
    if err != nil {
        return Data{}, fmt.Errorf("ошибка при разборе данных: %v", err)
    }
    return data, nil
}
func main() {
    fetchDataFromServer()
}