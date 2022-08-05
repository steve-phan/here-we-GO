# Go tutorial.

1. Deal with JSON file

- parse

```go
json.Unmarshal
```

- stringify(js)

```go
json.Marshal
```

2. Deal with environment variables

```go
    // load some ENV files first (".dev.env", prod.env)
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error, loading env fail")
	}
	// Load successfull then get Key
	key := os.Getenv("SUPER_KEY")
	log.Println("KEY", key)
```
