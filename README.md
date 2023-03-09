# flag4ctf

Flag build for CTFs. Provide a string and it build the "flag like" string.

## Usage

**Default**

```
flag4ctf make "Well played, you got the flag"

> FLAG{well_played_y0u_g0t_the_fl4g}
```

**Change the flag prefix**

```
flag4ctf make "Well played, you got the flag" -p "CTF"

> CTF{Well played, you got the flag}
```

**Change the flag case format**

```
flag4ctf make "Well played, you got the flag" -c lower

> FLAG{well_played_y0u_g0t_the_fl4g}
```

```
flag4ctf make "Well played, you got the flag" -c upper

> FLAG{WELL_PLAYED_Y0U_G0T_THE_FL4G}
```

## Easy Installation

### Binary Releases

We are now shipping binaries for each of the releases so that you don't even have to build them yourself! How wonderful is that!

If you're stupid enough to trust binaries that I've put together, you can download them from the [releases](https://github.com/LeoFVO/flag4ctf/releases) page.

### Using `go install`

If you have a [Go](https://golang.org/) environment ready to go (at least go 1.19), it's as easy as:

```bash
go install github.com/LeoFVO/flag4ctf@latest
```

PS: You need at least go 1.19 to compile flag4ctf.

### Using Docker

```bash
docker pull ghcr.io/leofvo/flag4ctf:latest
docker run flag4ctf:latest
```

### Building From Source

Since this tool is written in [Go](https://golang.org/) you need to install the Go language/compiler/etc. Full details of installation and set up can be found [on the Go language website](https://golang.org/doc/install). Once installed you have two options. You need at least go 1.19 to compile flag4ctf.

### Compiling

`flag4ctf` has external dependencies, and so they need to be pulled in first:

```bash
go get && go build
```

This will create a `flag4ctf` binary for you. If you want to install it in the `$GOPATH/bin` folder you can run:

```bash
go install
```

# License

See the LICENSE file.
