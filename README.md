####

Installation:

```
sudo apt install golang-go
go get -u github.com/markus-wa/demoinfocs-golang
```

Running:

You will need a demo file. You can find some demo files at [hltv.org](https://www.hltv.org/) like [this one](https://www.hltv.org/download/demo/20663)

The demo file should be in `/demos` and we will be passing your name as a argument.
Supposing that your demo has "mibr.dem" name:

```
go run main.go mibr
```

If the demo name is missing from the arguments, we have a default one: `demo`
In this case you need to rename your demo at `/demos` folder

For now, the demo name should be `mibr.dem`. It will be changed soon.
