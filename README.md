# UAParser
Command Line User Agent Parser

## Usage

```
$ uaparser 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_....'
1 "Mac OS X" "Google Chrome"

$ echo 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_....' | uaparser
```

```
$ uaparser --os 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_....'
1 "Mac OS X"
```

```
$ uaparser --browser 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_....'
1 "Google Chrome"
```

```
$ uaparser apache.log
100 "Mac OS X" "Google Chrome"
50 "Mac OS X" "Mozilla Firefox"
10 "Windows" "Google Chrome"
```

```
$ uaparser --os apache.log
150 "Mac OS X"
10 "Windows"
```
