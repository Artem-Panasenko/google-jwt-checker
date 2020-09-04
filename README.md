# google-jwt-checker

Parse and validate Google JWT without extra http-request to Google API.

## How to use

1. Set GOPRIVATE environment variable
```bash
export GOPRIVATE="github.rakops.com/SD"
```
2. Add ssh-keys to private github repository

3. And voila

```go

import (
    googlejwtchecker "github.rakops.com/sd/google-jwt-checker"
)
validToken, err := googlejwtchecker.Verify("TOKEN")

or 

validToken, err := googlejwtchecker.VerifyRequest(r *http.Request)


```


