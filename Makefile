run: build
	@./bin/app

build:
	@go build -o bin/app .

css:
	npx tailwindcss -i public/css/input.css -o public/css/style.css --watch   
