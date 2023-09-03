# FastAPI vs Gin

An experiment to compare performance of FastAPI and Gin.

## Overview:

-   [Setup](#setup)
-   [Run experiments](#run-experiments)

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
go run app.go
```

### wrk Benchmarking

Run the benchmarking script:

```bash
./benchmark.sh
```

**Note**: The script uses `wrk` to run benchmarks on http://localhost:8000.
You should benchmark one application at a time.
