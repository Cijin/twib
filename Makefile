run: build
	@./bin/app

build:
	@go build -o bin/app .

css:
	tailwindcss -i public/css/input.css -o public/css/style.css --watch   
