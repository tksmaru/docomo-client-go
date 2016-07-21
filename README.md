# docomo-client-go


## How to use

```go

    APIKey = "your API key value"
    d, err := NewDocomo(APIKey)
    if err != nil {
        fmt.Printf(err)
    }
    r, err := d.Dialogue("今日の天気はどうですか？")
    if err != nil {
        fmt.Printf(err)
    }
    fmt.Printf("response: %v", r)
```

## Testing

Test code requires API key as an environment variables. Set your API key like below.

```sh
export DOCOMO_API_KEY="xxxxxxxxxxxxxxxxxxxxx"
```

