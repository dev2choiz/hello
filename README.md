
# dev environment

## Requirements
- Create a `.env.local` file taking for example `.env.local.example`  
- Create a `./front/.env.local` file taking for example `./front/.env.local.example`

## Start the backend
```shell
make start
```
This command run in docker:
- 3 micro-services using **golang** and **gRPC** to expose the api
- The ESPv2 (Extensible Service Proxy) used as envoy sidecar
- Protoc container to generate protobuf files
- A postgres database

## Front
The front is done with **next.js** and **typescript**.  
It uses **grpc-web** to communicate with the gRPC api through the http2 protocol.  
This application is an example of how to do **SSR**, **SSG, **ISR** and **SPA** with next.js.

### Start the stack
_Go to the directory_
```shell
cd front
```
_Start the containers_
```shell
make docker-network up
```
_Install npm dependencies_
```shell
make install
```
_Generate the typescript protobufs from `.proto` files_
```shell
make gen-proto
```
_Start next app on http://localhost:3000_
```shell
make dev
```

# Deployment

## Micro-services
The micro-services are deployed in a GKE cluster provided by **terraform**.

A **jenkins** job triggers a gcp **cloud build** who will build a docker image in gcr and use **helm** to deploy pods in the cluster. 
The others resources (ingress, secret, ...) are generated with terraform.

## Front
The **next.js** front is deployed in google **cloud run**

As for micro-services, a jenkins job triggers a gcp cloud-build which builds a docker image then deploy it in a cloud run.
The load balancer and others resources are generated with terraform.
