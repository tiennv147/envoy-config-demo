# Envoy Config Demo
Here is the repository contains some usecases that I'm facing thought the time working with Envoy

## Categories
1. Circuit Breaking
2. Rate Limiting
3. Traffic Shifting
4. Header-based Routing
5. JWT Authentication
6. Service Mesh


## Demo for header-based routing
### Start docker-compose
```
cd 4-header-based-routing
docker-compose -f docker-compose.yml up
```
### Send request to test it
#### This reqeust will go to cluster-1 if the header x-cloud is present
```
curl --location --request GET 'http://127.0.0.1:8080/echo/v1/sample' \
--header 'x-cloud: echo-v1'
```
#### This request will go to cluster-2 if the is no header x-cloud or not match the value
```
curl --location --request GET 'http://127.0.0.1:8080/echo/v1/sample'
```