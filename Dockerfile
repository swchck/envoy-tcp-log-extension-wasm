FROM tinygo/tinygo:latest as build

WORKDIR /build
COPY . .
CMD "/bin/sh -c 'tinygo build -o log_ext.wasm -scheduler=none -target=wasi main.go'"
CMD "pwd"

FROM scratch

COPY --from=build /build/log_ext.wasm .