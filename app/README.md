# Dummy load api

This is a dummy api endpoint that is able to generate CPU and memory load for testing purposes. 
It allocates memory and splits the load across go routines (locked to an OS thread)

### Building locally

```bash
make build
```

Binary will be generated in the `build/` folder

### Building container

```bash
make build-image
```
To use container locally:

```bash
docker run -p 8080:8080 dummy-load-api
```

### Pushing the image to your preferred repo 
(Google Cloud Artifact Registry for example: https://cloud.google.com/artifact-registry)

```bash
docker tag dummy-load-api ${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPOSITORY_NAME}/dummy-load-api:latest
docker push ${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPOSITORY_NAME}/dummy-api:latest
```

# Kubernetes deployment

Modify k8s/dummy-load-api.yaml to your liking
```bash
kubectl apply -f k8s/dummy-load-api.yaml
```

# Usage / URL parameters

Example: 
```bash
curl http://localhost:8080/load?cores=2&time=1000&percentage=25&ram=1024
```

URL parameters are:

`cores` - amount of "cores" to use, will bind to an os thread - **default** 1

`time` - time in milliseconds - **default** 100(ms)

`percentage` - approximate cpu usage per core - **default** 25(%)

`ram` - amount of memory to allocate in megabytes (evenly distributes across threads) - **default** 128 (mb)
