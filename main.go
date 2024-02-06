package main

func main() {
	app := tgrid()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
