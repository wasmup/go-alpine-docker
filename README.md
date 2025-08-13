# go-alpine-docker
Create reduced binary size go compiler docker image

```sh
docker pull alpine:latest

wget https://go.dev/dl/go1.25.0.linux-amd64.tar.gz
tar -xzf go1.25.0.linux-amd64.tar.gz
rm -rf ./go/doc ./go/misc ./go/test ./go/api

docker build -t alpine-go:1.25.0 .
docker images --filter "dangling=true"
docker rmi $(docker images -q --filter "dangling=true")
docker images
# 197MB
docker history  alpine-go:1.25.0 

docker run --rm -it alpine-go:1.25.0 
go version
cat /etc/os-release || cat /usr/lib/os-release
exit

docker run --rm  alpine-go:1.25.0 go version
docker run --rm  alpine-go:1.25.0 uname -a

## App build
cd example-app
ls
docker build -t my-app:1.0.0 .

docker images --filter "dangling=true"
docker rmi $(docker images -q --filter "dangling=true")

docker images
# 14.4MB
docker history my-app:1.0.0
docker run --name my-app --network host --rm -d  my-app:1.0.0
docker ps
# firefox localhost:8080
curl -i http://localhost:8080
# docker logs -f my-app
docker stop my-app

docker rmi my-app:1.0.0

exit 

# Remove dangling build cache only 
docker builder prune	

# Remove dangling build cache without prompt 
docker builder prune -f	

# Remove all unused build cache 
docker builder prune -a	

# Remove unused containers, images, cache 
# docker system prune	

# Remove all unused containers, images (including dangling)
# docker system prune -a	

```
