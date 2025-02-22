dev: 
	make templ

# Run templ in watch mode with proxy
templ:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."

build-templ:
	templ generate

# Build everything
build: build-css build-templ