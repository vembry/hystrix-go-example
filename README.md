# hystrix-go-example

this is an example of hystrix-go usage

this example contains 2 service: 
1. `alpha` as our main service, circuit breaker are in place there
 - there are 2 api `/ping-a` and `ping-b`, both will do the same thing, with the only difference is,
 - `/ping-a` will not be using circuit breaker
 - `/ping-b` will be using circuit breaker

2. `zulu` as our secondary/dummy service for success and fail test
 - there is 1 api `/ping` that we will call from `alpha` service

the scenario is:
1. run service `alpha`
2. hit one of `alpha`'s endpoint
3. `alpha` will call `zulu`'s endpoint `/ping`
4. `zulu` return x to `alpha`
5. `alpha` return x to the requester

further explanation will be added