# docomo-client-go

[![CircleCI](https://circleci.com/gh/tksmaru/docomo-client-go.svg?style=svg&circle-token=584f14f264689884d04bd415118f82c725c5dcbc)](https://circleci.com/gh/tksmaru/docomo-client-go)
[![Coverage Status](https://coveralls.io/repos/github/tksmaru/docomo-client-go/badge.svg?branch=feature_ci)](https://coveralls.io/github/tksmaru/docomo-client-go?branch=feature_ci)


## Usage

### Import library
```go
import docomo "github.com/tksmaru/docomo-client-go"
```

### Dialogue
```go
    apiKey := "your API key value"
    c, err := docomo.NewClient(apiKey)
    if err != nil {
        fmt.Printf(err)
        return
    }
    r, err := c.Dialogue.Talk("今日の天気はどうですか？")
    if err != nil {
        fmt.Printf(err)
        return
    }
    fmt.Printf("response: %s", r.Utt)
```

### NamedEntity (As Individual user)
```go
    apiKey := "your API key value"
    c, err := docomo.NewClient(apiKey)
    if err != nil {
        fmt.Printf(err)
        return
    }
    r, err := c.NamedEntity.Extract("今日の5時の千葉の天気を千葉県庁の佐藤さんが確認した")
    if err != nil {
        fmt.Printf(err)
        return
    }
    fmt.Printf("response: %v", r)
```

### Configure http client
If you want to configure http client, initialize client like below.
```go
    apiKey := "your API key value"
    hc := &http.Client{}
    c, err := docomo.NewClient(apiKey, docomo.WithHttpClient(hc))
```

### Configure as corporation user
If you want to use API's for corporation account (ex: Named Entity), initialize client like below.
```go
    apiKey := "your API key value"
    c, err := docomo.NewClient(apiKey, docomo.AsCorp())
```


## Contribution

### Testing

Test code requires API key as an environment variables. Set your API key like below.

```sh
export DOCOMO_API_KEY="xxxxxxxxxxxxxxxxxxxxx"
```

