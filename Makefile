test: github.png github-full.png spiegel-light.png spiegel-dark.png spiegel-light-nojs.png spiegel-dark-nojs.png spiegel.pdf spiegel-nojs.pdf

clean:
	rm -v *.png *.pdf

github.png:
	go run ./pkg/main.go screenshot https://github.com github.png

github-full.png:
	go run ./pkg/main.go screenshot https://github.com github-full.png --wait 10s --full

github-footer.png:
	go run ./pkg/main.go screenshot https://github.com github-footer.png --wait 3s --selector=footer

spiegel-light.png:
	go run ./pkg/main.go screenshot https://spiegel.de spiegel-light.png --wait 3s

spiegel-dark.png:
	go run ./pkg/main.go screenshot https://spiegel.de spiegel-dark.png --darkmode --wait 3s

spiegel-light-nojs.png:
	go run ./pkg/main.go screenshot https://spiegel.de spiegel-light-nojs.png --javascript=false --wait 3s

spiegel-dark-nojs.png:
	go run ./pkg/main.go screenshot https://spiegel.de spiegel-dark-nojs.png --darkmode --javascript=false --wait 3s

spiegel.pdf:
	go run ./pkg/main.go pdf https://spiegel.de spiegel.pdf --wait 3s

spiegel-nojs.pdf:
	go run ./pkg/main.go pdf https://spiegel.de spiegel-nojs.pdf --javascript=false --wait 3s
