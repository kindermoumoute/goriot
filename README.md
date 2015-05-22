# goriot
--

#### Install :
	go get "github.com/kindermoumoute/goriot"

#### Using :
```go
	import "github.com/kindermoumoute/goriot"
```


### Example :
```go
	package main

	import "github.com/kindermoumoute/goriot"

	func main(){
		myapi := goriot.Get(APIkey, goriot.EUW, 10, 500)
	}
```