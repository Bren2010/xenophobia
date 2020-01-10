This repo just illustrates a bug I encountered when attempting to use `cloudflare/bn256` inside a golang `plugin`.

# Usage

To see what the problem is, do

```bash
cd cmd/unplugged
go build
./unplugged
```

This should write the number 1.

Now run

```bash
cd ../plugged/plug
go build -buildmode=plugin
cd ..
go build
./plugged
```

This should... well, it _should_ do exactly what the previous command did. What it does, is print a random number.

The only difference between the two calls is that the second loads the function using assembly as a plugin.

This occurs as of version 13.5 of go.
