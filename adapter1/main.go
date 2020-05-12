package main

func main() {
	var tee IAdaptee = new(Adaptee)
	var ad IAdapter = &Adapter{
		Ada: tee,
	}

	ad.Process(20)
}
