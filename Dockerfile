FROM tinygo/tinygo:latest as build

WORKDIR /build
COPY . .
RUN tinygo build -o log_ext.wasm -scheduler=none -target=wasi main.go

FROM scratch

COPY --from=build /build/log_ext.wasm /build/log_ext.wasm

WORKDIR /build
