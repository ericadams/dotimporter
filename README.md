# dotimporter
## Does using a dot import make unexported types or data directly accessible in Golang?
It doesn't look like it does:
```
| => go build
# github.com/ericadams/dotimporter
./main.go:6:2: imported and not used: "github.com/ericadams/dotimporter/inner"
./main.go:11:7: undefined: inner
./main.go:12:7: undefined: newPublicMember
```