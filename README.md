# echoserver

## Install

```bash
> kubectl apply -f https://raw.githubusercontent.com/mwennrich/echoserver/master/samples/echoserverStatefulSet.yaml
> kubectl apply -f https://raw.githubusercontent.com/mwennrich/echoserver/master/samples/echoserverService.yaml
```

## Use/Test

```bash
> curl http://example.com/headers
> curl http://example.com/hello
> curl http://example.com/echo
```
