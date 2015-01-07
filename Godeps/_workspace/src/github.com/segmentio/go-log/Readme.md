
# go-log

  Simple printf-style logger which is more or less the same as Go's core
  logger with log levels.

  View the [docs](http://godoc.org/github.com/segmentio/go-log).

## Example

```go
log.Debug("something")
log.Emergency("hello %s %s", "tobi", "ferret")

l := log.New(os.Stderr, log.DEBUG, "")
l.Debug("something happened")
l.Info("hello %s", "Tobi")
l.Error("boom something exploded")
```

## Conventions

 Typically you'll want to "inherit" from a parent logger, this allows
 for setting changes at the root of your application to trickle down.

 To do this you can use `Logger.New(prefix string)`, for example here
 a child logger is created by calling `.New()` on the standard logger:

```go
var log = log.Log.New("mylib")
```

## License

 MIT
