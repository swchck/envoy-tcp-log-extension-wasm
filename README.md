# Envoy TCP Logging Extension

## Build on Local Machine

```shell
# Run Go tests
docker compose run test
# Build *.wasm file and place into current dir
docker compose run build
```

## Store file inside container

```shell
docker build -t envoy-wasm-ext-tcp-log .
```

## Configuration Example _(envoy.yml)_

```yaml
static_resources:
  listeners:
    - name: main
      address:
        socket_address:
          address: 0.0.0.0
          port_value: 18000
      filter_chains:
        # Create new WASM filter
        - filters:
            - name: envoy.filters.network.wasm
              typed_config:
                "@type": type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
                config:
                  vm_config:
                    runtime: "envoy.wasm.runtime.v8"
                    code:
                      local:
                        # Place *.wasm filename below
                        filename: "log_ext.wasm"
```

## FYI
- [POSSIBLE LOG VARIABLES](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/advanced/attributes)
- [PLUGIN LIBRARY OVERVIEW](https://github.com/tetratelabs/proxy-wasm-go-sdk/blob/9c66f8cb17b6f3b7e3baba4302881891137b7163/doc/OVERVIEW.md)