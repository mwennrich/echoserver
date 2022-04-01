# echoserver

## Install

```bash
> kubectl apply -f https://raw.githubusercontent.com/mwennrich/echoserver/main/samples/echoserverStatefulSet.yaml
> kubectl apply -f https://raw.githubusercontent.com/mwennrich/echoserver/main/samples/echoserverService.yaml
```

## Use/Test

```bash
> curl http://example.com/headers
> curl http://example.com/hello
> curl http://example.com/stream
> curl http://example.com/stream?interval=0.5s
> echo -n "blafasel"| curl  --data-urlencode data@- http://example.com/echo
```
