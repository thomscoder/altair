package build

import "altair/src/view"


var articles = []view.Article{
{Title: `docs`, Author: `Thomas`, Body: ``, Path: `docs`, IsDir: true}, {Title: `gcd`, Author: `Thomas`, Body: ``, Path: `docs/gcd`, IsDir: true}, {Title: `gcd.wat`, Author: `Thomas`, Body: `(module
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
`, Path: `docs/gcd/gcd.wat`, IsDir: false}, {Title: `hello`, Author: `Thomas`, Body: ``, Path: `docs/hello`, IsDir: true}, {Title: `hello.wat`, Author: `Thomas`, Body: `;; An example of instantiating a small wasm module 
;; which imports functionality from the host, 
;; then calling into wasm which calls back into the host.

(module
    (import "" "hello" (func $hello))
    (func (export "run")
    (call $hello)
    )
)`, Path: `docs/hello/hello.wat`, IsDir: false}, }
