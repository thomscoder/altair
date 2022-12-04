package build

import "altair/src/view"


var articles = []view.Article{
{Title: `docs`, Author: `Thomas ✌️`, Body: ``, Path: `docs`, IsDir: true}, {Title: `hello`, Author: `Thomas ✌️`, Body: ``, Path: `docs/hello`, IsDir: true}, {Title: `helloWorld.txt`, Author: `Thomas ✌️`, Body: `You can write normal text like this

Colored text like [blue]this[white]

You can import scripts like this:
[orange]
(module
    (import "" "hello" (func $hello))
    (func (export "run")
    (call $hello)
    )
)
[white]

(this is limited to one external script per page for now)`, Path: `docs/hello/helloWorld.txt`, IsDir: false, Snippet: `(module
    (import "" "hello" (func $hello))
    (func (export "run")
    (call $hello)
    )
)`}, {Title: `welcome.txt`, Author: `Thomas ✌️`, Body: `[red]Welcome on Altair ⭐

Generate portable terminal based documentation.

Please refer to the README for more information`, Path: `docs/welcome.txt`, IsDir: false,}, }
