build:
	oapi-codegen -package tcgplayer -generate=types,client openapi.yml > client.gen.go
	go build
