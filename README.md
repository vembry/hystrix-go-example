# hystrix-go-example

this is an example of [hystrix-go](https://github.com/afex/hystrix-go) usage in web dev using gin

## Explanation

this example contains 2 service: 
1. `alpha` as our main service, circuit breaker are in place there
   - there are 2 api `/ping-a` and `ping-b`, both will do the same thing, with the only difference is,
   - `/ping-a` will not be using circuit breaker
   - `/ping-b` will be using circuit breaker

2. `zulu` as our secondary/dummy service for external service
   - there is 1 api `/ping` that we will use as dummy endpoint

## Scenario
1. run service `alpha`
2. hit one of `alpha`'s endpoint
3. `alpha` will call `zulu`'s endpoint `/ping`
4. `zulu` return x to `alpha`
5. `alpha` return x to the requester

## How to run:
1. you need to run both separately
2. run alpha from root folder
 ```bash
 cd alpha; go run main.go
 ```
3. run zulu from root folder
 ```bash
 cd zulu; go run main.go
 ```

#

further explanation will be added
