# FastAPI vs Gin vs Fiber

An experiment to compare performance of FastAPI, Gin and Fiber.

## Overview:

-   [Setup](#setup)
-   [Run experiments](#run-experiments)
-   [Results](#results)

## Setup

### FastAPI

The repository uses Python 3.11. Follow the steps below to get started:

-   Create conda environment:

    ```bash
    conda create -n fastapi-tests python=3.11 -y
    conda activate fastapi-tests
    ```

    You can choose any other environment manager of your choice.

-   Install dependencies:

    ```bash
    pip install -r requirements.txt
    ```

    **Note**: All requirement files are generated using `pip-tools`.

### Gin

The repository uses Go 1.21. Follow the steps below to get started:

-   Install dependencies:

    ```bash
    go mod tidy
    ```

## Run experiments

Both FastAPI and Gin applications serve one endpoint:

- `/`: sleeps for 10 seconds and returns a response

### FastAPI

The experiment uses `gunicorn` server to run the FastAPI application.

```bash
gunicorn -k uvicorn.workers.UvicornWorker app:app
```

### Gin

The experiment uses `gin` server to run the Gin application.

```bash
go run gin/app.go
```

### Fiber

The experiment uses `fiber` server to run the Fiber application.

```bash
go run fiber/app.go
```

### wrk Benchmarking

Run the benchmarking script:

```bash
./benchmark.sh
```

**Note**: The script uses `wrk` to run benchmarks on http://localhost:8000.
You should benchmark one application at a time.

## Results

### FastAPI

```bash
Running 1m test @ http://localhost:8000/
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    10.01s     5.65ms  10.02s    83.40%
    Req/Sec    23.08     52.04   180.00     87.50%
  500 requests in 1.00m, 71.29KB read
Requests/sec:      8.33
Transfer/sec:      1.19KB
```

### Gin

```bash
Running 1m test @ http://localhost:8000/
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    10.01s     4.34ms  10.02s    57.20%
    Req/Sec     9.87     21.25    70.00     86.96%
  500 requests in 1.00m, 70.31KB read
Requests/sec:      8.33
Transfer/sec:      1.17KB
```

### Fiber

```bash
Running 1m test @ http://localhost:8000/
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    10.00s     3.22ms  10.01s    80.00%
    Req/Sec     2.00      0.00     2.00    100.00%
  500 requests in 1.00m, 62.99KB read
Requests/sec:      8.33
Transfer/sec:      1.05KB
```
