package main

func main() {
	phone := Phone{
		V5V: &VoltageAdapter{
			V220V: &Votagle220V{},
		},
	}

	phone.Charge()
}
