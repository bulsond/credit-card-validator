package main

func main() {
	app, err := NewApp("banks.txt", 13, 19)
	if err != nil {
		panic(err)
	}

	app.Run()
}
