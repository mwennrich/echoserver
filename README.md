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
> curl http://example.com/echo
> curl http://example.com/echo?interval=0.5
```
