build: generate
	go build
generate:
	oapi-codegen -package tcgplayer -generate=types,client openapi.yml > client.gen.go

