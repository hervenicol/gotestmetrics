# Test go app metrics

## Go app

Retrieves random jokes from https://icanhazdadjoke.com/

### endpoints tests

Simple call for 1 joke:

```
curl -s localhost:8080/joke | jq '.'
```

Call for 3 jokes:
```
curl -s -X GET -d '{"qty":3}' localhost:8080/joke | jq '.'
```

Ask for too many jokes:
```
curl -s -X GET -d '{"qty":14}' localhost:8080/joke | jq '.'
```

Check metrics:
```
curl -s localhost:8080/metrics
```

## Deployment

All is in `helm/` subdirectory.
