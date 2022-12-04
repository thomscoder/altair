# Altair ⭐

Generate portable terminal based documentation.
<br/>
## How to use

I will soon turn it into a module so it will ease the generation
For now you can simply `clone` this repository and modify the content of the `./docs` folder.

Altair supports nested folders.

<img src="https://i.ibb.co/gM2ZGyy/Screenshot-2022-12-04-alle-17-14-02.png" alt="Screenshot-2022-12-04-alle-17-14-02" border="0">

Run:

```bash
pnpm dev
```
or

```bash
pnpm build
```

The `building process` is simple. 

- It first walks the `./docs` directory 
- Generates a new file with all the computed content (like a db)
- Builds

Altair builds the tree from the generated content (and not from the `os.Readir` or `os.Walkdir` or whatnot)

---

<strong>Write your stuff in `.txt` </strong>

- Add colors: altair is based on Tview so it supports its coloring syntax `[COLOR]something[RESET]`<br/>
  for example if I wanted to color `Hello world` in orange I'd write `[orange]Hello world[white]`

- Add snippets: (for now only one snippet per `.txt`) simply wrap the file name, that contains the snippet, within `-` signs<br/>
  for example if I wanted to import snippet `hello.wat`, I'd write <br/>

  <pre>I'm now importing hello.wat content

    \-
    hello.wat
    \-
    done!
  </pre>

  Altair will replace it with the file content

  > IMPORTANT: files must be in the same dir (working on it)

<strong>Quit:</strong><br/>
Press `q` to quit.
<br/>

## Light/Dark theme

There's the possibility to switch between light and dark theme (by clicking the buttons)

<img src="https://i.ibb.co/3MQRvzR/Screenshot-2022-12-04-alle-17-13-48.png" alt="Screenshot-2022-12-04-alle-17-13-48" border="0">
<br/>

## Roadmap

As for now it is still an MVP

- Search (search in the docs)
- Simplify the process
- Run snippets

