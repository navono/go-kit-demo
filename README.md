# gokit-demo
## stringsvc
launch three instances:
> go run stringsvc/*.go -listen=:8085 &
<br>
> go run stringsvc/*.go -listen=:8086 &
<br>
> go run stringsvc/*.go -listen=:8087 &

then add another instance with proxy:
> go run stringsvc/*.go-listen=:8088 -proxy=localhost:8085,localhost:8086,localhost:8087

test:
```
for s in foo bar baz ; do curl -d "{\"s\":\"$s\"}" localhost:8086/uppercase ; done
```