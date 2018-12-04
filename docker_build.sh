docker build -t biru-adaptor -f Dockerfile.biru-adaptor .
docker save biru-adaptor > biru-adaptor.tar
docker build -t biru-rest -f Dockerfile.biru-rest .
docker save biru-rest > biru-rest.tar
docker build -t blanja-rest -f Dockerfile.blanja-rest .
docker save blanja-rest > blanja-rest.tar
docker build -t blanja-adaptor -f Dockerfile.blanja-adaptor .
docker save blanja-adaptor > blanja-adaptor.tar