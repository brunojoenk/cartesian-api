## Cartesian-api
This app deal with a series of points represented as (x,y) coordinates on a simple 2-dimensional plane using the Manhattan distance method.
This app is running on Heroku Clound at address: https://cartesian-api.herokuapp.com/

### Requiriments
- Golang 1.17+

### Environment vars:
```
- PORT (Default is 8080)
- PATH_DATA_POINTS (Default is 'data/points.json')
```

### How to run locally

You can use Makefile to use commands of the app:

To build and run project:
```
make build-run
```

To run main.go:
```
make run-main
```

To run unit tests:
```
make tests
```

To run coverage tests:
```
make tests-coverage
```

### APIs
#### To calculate points
```
/api/points?x={integerParam}&y={integerParam}&distance={integerParam}
```

#### An example: 
```
https://cartesian-api.herokuapp.com/api/points?x=1&y=1&distance=20 
```
#### or when run locally:
```
http://localhost:8080/api/points?x=1&y=1&distance=20
```
