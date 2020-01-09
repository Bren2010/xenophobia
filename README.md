This repo just illustrates a bug I encountered when attempting to use `cloudflare/bn256` inside a golang `plugin`.

I suspect this has something to do with the usage of assembly in the `bn256` repo, but I don't know enough to verify it.

# Usage

To see what the problem is, do

```bash
cd cmd/unplugged
go build
./unplugged
```

This should write the same sequence of bytes two times, and it means it successfully generated, marshalled and unmarshalled an element of `bn256.G2`.

Now run

```bash
cd ../plugged/plug
go build -buildmode=plugin
cd ..
go build
./plugged
```

This should... well, it _should_ do exactly what the previous command did. What it does, is report an error "malformed point", which essentially means the resulting point is not on the curve.

The only difference between the two calls is that the second loads the function using `bn256` as a plugin.

This occurs on version 13.5 of go.
