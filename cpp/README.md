# Simple HTTP/1.1 server

Implementation of a simple HTTP server in C++

## Quick start

```bash
mkdir build && cd build
cmake ..
make
./test_SimpleHttpServer # Run unit tests
./SimpleHttpServer      # Start the HTTP server on port 8080
```

- There are two endpoints available at `/` and `/hello.html` which are created for demo purpose.
- In order to have multiple concurrent connections, make sure to raise the resource limit (with `ulimit`) before running the server. A non-root user by default can have about 1000 file descriptors opened, which corresponds to 1000 active clients.