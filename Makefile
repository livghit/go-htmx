run:
	go run main.go

tailwind:
	cd client && npx tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch
