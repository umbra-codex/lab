package main

func main() {
	config := &Config{
		Commands: getCommands(),
	}
	startRepl(config)
}
