package main

func main() {
	var btn IButton = &MyBtnAdapter{}
	btn.Click()
}
