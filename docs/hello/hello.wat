;; An example of instantiating a small wasm module 
;; which imports functionality from the host, 
;; then calling into wasm which calls back into the host.

(module
    (import "" "hello" (func $hello))
    (func (export "run")
    (call $hello)
    )
)