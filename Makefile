test: github.png spiegel.png spiegel-darkmode.png spiegel.pdf

clean:
	rm -v *.png *.pdf

github.png:
	go run ./pkg/main.go screenshot https://github.com github.png --wait 3s

spiegel.png:
	go run ./pkg/main.go screenshot https://spiegel.de spiegel.png --wait 3s

spiegel-darkmode.png:
	go run ./pkg/main.go screenshot https://spiegel.de spiegel-darkmode.png --darkmode --wait 3s

spiegel.pdf:
	go run ./pkg/main.go pdf https://spiegel.de spiegel.pdf --wait 3s
