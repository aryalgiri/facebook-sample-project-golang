## Sample FB Service 

### Steps to run the service

#### Requires:

* ![Install Golang](https://golang.org/doc/install)

* ![Install Golang CI Lint](https://github.com/golangci/golangci-lint)

* ![Install Go Dep](https://github.com/golang/dep)


#### How to run

##### Webserver

## Build Docker container and tag the container with name
```
docker build -t sample-fb .
```

## Run the container 
```
docker run --publish 8080:8080 --name sample-fb --rm sample-fb
```

## To see the results
```
localhost:8080
```

##### To check the code Quality Golang CI Lint

```
golangci-lint run
```
