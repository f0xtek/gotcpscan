# gotcpscan

Concurrent TCP Scanner.

## Usage

```
$ gotcpscan -h
Usage of ./gotcpscan:
  -host string
        The host you wish to scan. Default: 127.0.0.1 (default "127.0.0.1")
  -ports string
        The ports you wish to scan. Accepts nmap formatted port specifications. Default: 1-65535 (default "1-65535")
```

### Examples

```
$ gotcpscan -host scanme.nmap.org
22 open
80 open
9929 open
31337 open
```

```
$ gotcpscan -host scanme.nmap.org -ports 80
80 open
```

```
$ gotcpscan -host scanme.nmap.org -ports 22,80
22 open
80 open
```

```
$ gotcpscan -host scanme.nmap.org -ports 1-1000
22 open
80 open
```

```
$ gotcpscan -host scanme.nmap.org -ports 1-1000,31337
22 open
80 open
31337 open
```

## Installing

`go get -u github.com/f0xtek/gotcpscan`
