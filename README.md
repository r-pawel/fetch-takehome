# fetch-takehome

## Running the program

### Docker commands to build and run the program to listen to port 8080

```shell
docker build -t r-pawel/fetch-takehome:latest .   
docker run -p 8080:8080 r-pawel/fetch-takehome:latest
```
### Run without docker
```shell
go run cmd/server/main.go
```

## Assumptions
* Only stored how many points a receipt is worth
* For time points didn't include 2pm and 4pm on the dot
* Not thread safe 
* Individual item costs must add up to the total on the receipt to be valid