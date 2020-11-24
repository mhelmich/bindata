bindata:
	cd cmd/bindata && go build
	./cmd/bindata/bindata -o templates_bindata.go -package bindata templates/*.tmpl
.PHONY: bindata
