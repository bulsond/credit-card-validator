package main

func main() {
	app, err := NewApp("banks.txt")
	if err != nil {
		panic(err)
	}

	app.Run()
}
