# url-json



## Installation

```shell
go install github.com/raahii/url-json@v1.0.1
```

Or download the binary from [GitHub Releases](https://github.com/raahii/url-json/releases) and drop it in your `$PATH`.


## Usage

```shell
$ url-json 'https://user:pass@example.com:1234/path1/path2/?q1=v1&q2=v2-1&q2=v2-2#frag' | jq
{
  "scheme": "https",
  "user": {
    "username": "user",
    "password": "pass"
  },
  "host": "example.com",
  "port": "1234",
  "path": "/path1/path2/",
  "fragment": "frag",
  "queries": {
    "q1": "v1",
    "q2": [
      "v2-1",
      "v2-2"
    ]
  }
}
```
