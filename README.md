# Go api project template
This is a generic template for making my future go projects. Please feel free to
add or modify this template if you think it can be made better. I am trying to
make it a slim framework type which can work as soon as it is cloned.

## Installation
Please clone this repository and then go to go.mod and replace `appname` from
entire project to the project name of your choice, then run
```shell
go get ./...
go mod vendor
```
And finally run
```shell
docker-compose up -d 
```

### TODO
- [ ] I still have to add a way to hot reload the application
- [ ] Better build tools using make
- [ ] Have to add a sh toolkit that will do the installation automatically
