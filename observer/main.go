package main

import "fmt"

func main() {
	sina := &Sina{}
	baidu := &Baidu{}
	wd := WeatherData{}
	wInfo := &WeatherInfo{
		Temperature: 36,
		Pressure:    150,
		Humidity:    40,
	}
	wd.Add(sina)
	wd.Add(baidu)
	fmt.Println("开始通知所有观察者======")
	wd.SetData(wInfo)

	wd.Remove(sina)
	fmt.Println("开始通知所有观察者======")
	wd.SetData(wInfo)
}
