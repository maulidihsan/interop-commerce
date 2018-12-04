CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -i -v -o bin/biru-adaptor/server github.com/maulidihsan/interop-commerce/cmd/biru-adaptor
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -i -v -o bin/biru-rest/server github.com/maulidihsan/interop-commerce/cmd/biru-rest
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -i -v -o bin/blanja-adaptor/server github.com/maulidihsan/interop-commerce/cmd/blanja-adaptor
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -i -v -o bin/blanja-rest/server github.com/maulidihsan/interop-commerce/cmd/blanja-rest
