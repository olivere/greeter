# Tiny microservices example

This is a tiny microservices example done with [micro](https://github.com/micro).

It illustrates setting up a simple service, running several copies of it as
as server, and calling the service as a client.

## Prerequisites

1. Install Go (see [here](https://golang.org/)).
2. Install Consul (see [here](https://www.consul.io/intro/getting-started/install.html))
3. Run Consul (`consul agent -dev -advertise=127.0.0.1`)

## Run the example

1. Go get this repo:

```sh
$ go get github.com/olivere/greeter
```

2. Install the protobuf generators for micro:

```sh
$ go get github.com/micro/protobuf/{proto,protoc-gen-go}
```

3. Build the executable

```sh
make
```

4. Run the server:

```sh
./greeter
```

5. Open a 2nd terminal window and run the client:

```sh
./greeter -client -name=Oliver
```

## Illustrating several copies of the server

1. Open 3 terminal windows and run a copy of the greeter service:

```sh
$ ./greeter -id=A
$ ./greeter -id=B
$ ./greeter -id=C
```

2. Start a client in an infinite loop:

```sh
$ while true; do ./greeter -client; sleep 1; done
```

Now start and stop services as you like, and see how the client reacts.

# License

MIT
