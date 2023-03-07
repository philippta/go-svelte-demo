# go-svelte-demo

Demo on how to embed svelte components into dynamically rendered Go templates.

### Main ingredients

1. Go's `html/template` for rendering
2. node packages:
    - esbuild
    - esbuild-svelte

### How it's built

1. Create an ordinary Svelte component (.svelte) in `templates/components`.
2. Include that component in `templates/components/index.js` and attach it to the JS window object.
3. Bundle the components using `node build.js` (output lands in `/static/components.js`)
4. Include that JS file in your html
5. Embed the component in your dynamically rendered HTML like so:
```html
<div id="my-search-table"></div>
<script>
    new SearchTable({
        target: document.getElementById("my-search-table"),
        props: {
            data: {{ json . }},
        },
    })
</script>
```

### How to run this demo

1. `npm install`
2. `node build.js`
3. `go run main.go`
4. Visit [http://localhost:8080](localhost:8080)
