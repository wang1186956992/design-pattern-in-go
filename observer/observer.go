package main

import "fmt"

type WeatherInfo struct {
	Temperature float64
	Pressure    float64
	Humidity    float64
}

type IObserver interface {
	Update(info *WeatherInfo)
}

type Baidu struct {
	WI *WeatherInfo
}

func (o *Baidu) Update(info *WeatherInfo) {
	o.WI = info
	o.display(o.WI)
}

func (o *Baidu) display(info *WeatherInfo) {
	fmt.Println("百度天气预报======")
	fmt.Printf("温度 %v\n", info.Temperature)
	fmt.Printf("气压 %v\n", info.Pressure)
	fmt.Printf("湿度 %v\n", info.Humidity)
}

type Sina struct {
	WI *WeatherInfo
}

func (s *Sina) Update(info *WeatherInfo) {
	s.WI = info
	s.display(s.WI)
}

func (s *Sina) display(info *WeatherInfo) {
	fmt.Println("新浪天气预报======")
	fmt.Printf("温度 %v\n", info.Temperature)
	fmt.Printf("气压 %v\n", info.Pressure)
	fmt.Printf("湿度 %v\n", info.Humidity)
}

type ISubject interface {
	Add(o IObserver)
	Remove(o IObserver)
	NotifyAll()
}

type WeatherData struct {
	WI        *WeatherInfo
	observers []IObserver
}

func (w *WeatherData) Add(o IObserver) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) Remove(o IObserver) {
	targetIndex := 0
	for idx, ele := range w.observers {
		if ele == o {
			targetIndex = idx
			break
		}
	}
	left := w.observers[:targetIndex]
	right := w.observers[targetIndex+1:]
	w.observers = append(left, right...)
}

func (w *WeatherData) SetData(info *WeatherInfo) {
	w.WI = info
	w.NotifyAll()
}

func (w *WeatherData) NotifyAll() {
	for _, o := range w.observers {
		o.Update(w.WI)
	}
}
