## Oxpecker Plugins

This repo holds a set of default plugins for oxpecker to work with Go and the Buffalo stack used by the Wawandco team. 
### Usage

You can use individual plugins in here or use the base set of plugins in the `All` variable.

```go
// in cmd/ox/main.go
import (
    ...
    "github.com/wawandco/oxplugins"
)

func main() {
    fmt.Print("~~~~ Using {{.Name}}/cmd/ox ~~~\n\n")
    ctx := context.Background()

    pwd, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    
    cl := cli.New()
    // append your plugins here
    cl.Plugins = append(cl.Plugins, oxplugins.Base...)
    ...
}
```


