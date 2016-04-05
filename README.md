# Tiny microservice example

This is a tiny microservice example done with [micro](https://github.com/micro).
It illustrates setting up a simple service, running several copies, and
calling the service as a client.

## Prerequisites

1. Install Go (see [here](https://golang.org/)).
2. Install Consul (see [here](https://www.consul.io/intro/getting-started/install.html))
3. Run Consul (`consul agent -dev -advertise=127.0.0.1`)

## Run the example

Step 1. Go get this repo:

```sh
$ go get github.com/olivere/greeter
```

Step 2. Install the protobuf generators for micro:

```sh
$ go get github.com/micro/protobuf/{proto,protoc-gen-go}
```

Step 3. Build the executable

```sh
make
```

Step 4. Run the server:

```sh
./greeter
```

Step 5. Open a 2nd terminal window and run the client:

```sh
./greeter -client -name=Oliver
```

You should now see `Hello Oliver` printed on the screen.

## Illustrating several copies of the server

Step 1. Open 3 terminal windows and run a copy of the greeter service:

```sh
$ ./greeter -id=A
$ ./greeter -id=B
$ ./greeter -id=C
```

Step 2. Start a client in an infinite loop:

```sh
$ while true; do ./greeter -client; sleep 1; done
```

You should now see replies from a random service (A, B, or C) as micro uses
a random selector by default.

Now experiment with it and e.g. start and stop services, and see how the client reacts.

# License

MIT
