# log4go
log4go

```go get github.com/flywithbug/log4go```

```
func SetLog() {
	w := log.NewFileWriter()
	w.SetPathPattern("./log/log-%Y%M%D.log")
	c := log.NewConsoleWriter()
	c.SetColor(true)
	log.Register(w)
	log.Register(c)
	log.SetLevel(log.DEBUG)
	log.SetLayout("2006-01-02 15:04:05")
}
```

```		
log4go.Warn(err.Error())
log4go.Info("info need log")

```
![log4go](http://7xqjzu.com1.z0.glb.clouddn.com/1517409614395.jpg)

