# Nexmo Server

A microservice for Nexmo SMS API

## Getting started

1. Install Consul

  Consul is the default registry/discovery for go-micro apps. It's however pluggable.
  [https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

2. Run Consul
  ```
  $ consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul --advertise=127.0.0.1
  ```

3. Get your API keys from Nexmo.

4. Download and start the service

  ```shell
  go get github.com/jorjeb/nexmo-srv
  nexmo-srv --nexmo_api_key=XXX --nexmo_api_secret=XXX --registry_address=127.0.0.1
  ```

  OR as a docker container

  ```shell
  nexmo-srv --nexmo_api_key=XXX --nexmo_api_secret=XXX --registry_address=127.0.0.1
  ```

## The API

```shell
micro query go.micro.srv.nexmo SMS.Send '{"to":"09XXXXXXXXX","from":"Test","regionCode":"PH","text":"this is a test"}'
```
