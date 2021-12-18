# jsonquery

jsonquery is a utility for querying JSON data.

## Installation

```bash
go get github.com/Yz4230/jsonquery
```

## Usage

### Prepare data

```go
var jsonString = []byte(`{
  "name": "John Doe",
  "age": 42,
  "children": [
    {
      "name": "Alice",
      "age": 13
    },
    {
      "name": "Bob",
      "age": 12
    }
  ],
}`)

var d interface{}
json.Unmarshal(jsonString, &d)
```

### Simple key

```go
v, _ := jsonquery.New(d).Key("name").End()
fmt.Println(v) // John Doe

v, _ = jsonquery.New(d).Key("age").End()
fmt.Println(v) // 42

v, _ = jsonquery.New(d).Key("children").End()
fmt.Println(v) // [{Alice 13} {Bob 12}]
```

### Nested key

```go
v, _ = jsonquery.New(d).Key("children").Key("name").End()
fmt.Println(v) // [Alice Bob]
```
