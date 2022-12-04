package build

import "altair/src/view"


var articles = []view.Article{
{Title: `docs`, Author: `Thomas ✌️`, Body: ``, Path: `docs`, IsDir: true}, {Title: `gcd`, Author: `Thomas ✌️`, Body: ``, Path: `docs/gcd`, IsDir: true}, {Title: `example.txt`, Author: `Thomas ✌️`, Body: `Hello from the example
This example takes the gcd.wat script

[orange]
(module
  (func $gcd (param i32 i32) (result i32)
    (local i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        br_if 0 (;@2;)
        local.get 1
        local.set 2
        br 1 (;@1;)
      end
      loop  ;; label = @2
        local.get 1
        local.get 0
        local.tee 2
        i32.rem_u
        local.set 0
        local.get 2
        local.set 1
        local.get 0
        br_if 0 (;@2;)
      end
    end
    local.get 2
  )
  (export "gcd" (func $gcd))
)
[white]

and runs it`, Path: `docs/gcd/example.txt`, IsDir: false, Snippet: `(module
  (func $gcd (param i32 i32) (result i32)
    (local i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        br_if 0 (;@2;)
        local.get 1
        local.set 2
        br 1 (;@1;)
      end
      loop  ;; label = @2
        local.get 1
        local.get 0
        local.tee 2
        i32.rem_u
        local.set 0
        local.get 2
        local.set 1
        local.get 0
        br_if 0 (;@2;)
      end
    end
    local.get 2
  )
  (export "gcd" (func $gcd))
)`}, {Title: `hello`, Author: `Thomas ✌️`, Body: ``, Path: `docs/hello`, IsDir: true}, {Title: `helloWorld.txt`, Author: `Thomas ✌️`, Body: `A simple Hello world program in WASM

An example of instantiating a small wasm module which imports functionality from the host, 
then calling into wasm which calls back into the host.

[orange]
(module
    (import "" "hello" (func $hello))
    (func (export "run")
    (call $hello)
    )
)
[white]

This prints "hello world" in the console`, Path: `docs/hello/helloWorld.txt`, IsDir: false, Snippet: `(module
    (import "" "hello" (func $hello))
    (func (export "run")
    (call $hello)
    )
)`}, {Title: `welcome.txt`, Author: `Thomas ✌️`, Body: `[red]A portable WebAssembly guide`, Path: `docs/welcome.txt`, IsDir: false,}, }
