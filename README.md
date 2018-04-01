# moe

simple cli spinner by go.

## Usage

```bash
go get -u github.com/biezhi/moe
```

```go
moe := moe.New("正在为您加载").Start()
time.Sleep(4 * time.Second)
moe.Stop()
```

## License

[MIT](LICENSE)