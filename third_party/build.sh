CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -i -v -o bin/blibli-adaptor/server github.com/maulidihsan/flashdeal-webservice/cmd/blibli-adaptor
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -i -v -o bin/blibli-rest/server github.com/maulidihsan/flashdeal-webservice/cmd/blibli-rest
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -i -v -o bin/blanja-adaptor/server github.com/maulidihsan/flashdeal-webservice/cmd/blanja-adaptor
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -i -v -o bin/blanja-rest/server github.com/maulidihsan/flashdeal-webservice/cmd/blanja-rest
