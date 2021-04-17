# musicflow

Speaks the LG Music Flow protocol. This package and its commands can be used to control LG speakers that are controllable via the Music Flow app.

The API is mostly implemented, more can be added as needed.

## Usage

As a module.

```console
go get -u github.com/mafredri/musicflow
```

Tool for controlling the speakers.

```console
go get -u github.com/mafredri/musicflow/cmd/mufloctl
mufloctl -addr soundbar.local -nightmode=false
```

Run as wasm (node):

```console
env GOOS=js GOARCH=wasm go build github.com/mafredri/musicflow/cmd/mufloctl -o mufloctl.wasm

/usr/local/opt/go/libexec/misc/wasm/go_js_wasm_exec mufloctl.wasm -addr soundbar.local
```
