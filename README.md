# docomo-client-go


## Usage

### Import library
```go
import docomo "github.com/tksmaru/docomo-client-go"
```

### Dialogue

```go

    apiKey := "your API key value"
    d, err := docomo.NewDialogue(apiKey)
    if err != nil {
        fmt.Printf(err)
        return
    }
    r, err := d.Talk("今日の天気はどうですか？")
    if err != nil {
        fmt.Printf(err)
        return
    }
    fmt.Printf("response: %s", r.Utt)

```

## Contribution

### Testing

Test code requires API key as an environment variables. Set your API key like below.

```sh
export DOCOMO_DIALOGUE_API_KEY="xxxxxxxxxxxxxxxxxxxxx"
```

