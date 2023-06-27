# WASM Talk

Package wasmtalk contains helper types and functions for writting go based scripts and hooks for [CartJumper](https://cartjumper.com), and [FreightChick](https://freightchick.com).

## Protocol

GC Applications utilize a JSON message based protocol to pass messages back and forth over stdout and stdin between webassembly scripts and our applications.

Create a WebAssembly module using the WASI compilation mode in golang and write messages to stdout to return data from your scripts.

```
[GC Application] - {JSON Message} -> [Script/Hook]
[GC Application] <- {JSON Message} - [Script/Hook]
...

```
