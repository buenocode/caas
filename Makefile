test: github.png spiegel.png spiegel-darkmode.png

github.png:
	go run ./pkg/main.go screenshot https://github.com github.png

spiegel.png:
	go run ./pkg/main.go screenshot https://spiegel.de spiegel.png

spiegel-darkmode.png:
	go run ./pkg/main.go screenshot https://spiegel.de spiegel-darkmode.png --darkmode
