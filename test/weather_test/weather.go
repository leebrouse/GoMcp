package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CurrentWeather struct {
	Main struct {
		Temp     float32 `json:"temp"`     // 温度
		Humidity float32 `json:"humidity"` // 湿度
		Pressure float32 `json:"pressure"` // 气压
	} `json:"main"`

	Wind struct {
		Speed float32 `json:"speed"` // 风速
	} `json:"wind"`
}

func main() {
	// 用你的API Key替换下面的字符串
	// apiKey := "fabbc7c41b53a66cf4d3a514a135a015"
	// location := "shanghai"

	// 构造API URL
	url := "https://api.openweathermap.org/data/2.5/weather?q=shanghai,cn&APPID=5cbca3a01d81fdc6d4428c102425fe70&units=metric"

	// 发送HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 解析JSON数据
	var weatherData CurrentWeather
	fmt.Println(string(body))
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Printf("weatherData: %.2f °C\n", weatherData)

	// 打印获取到的数据
	fmt.Printf("Temperature: %.1f °C\n", weatherData.Main.Temp)
	fmt.Printf("Humidity: %.2f %%\n", weatherData.Main.Humidity)
	fmt.Printf("Pressure: %.2f hPa\n", weatherData.Main.Pressure)
	fmt.Printf("Wind Speed: %.1f m/s\n", weatherData.Wind.Speed)
}
