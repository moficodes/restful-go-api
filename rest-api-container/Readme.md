# Application Delivery

The best application is a running application. For our we worked on creating some amazing REST API so far but its no use to anyone if we don't have it accessible to others.

To have a Go REST API deployed we have a number of options. These are a few of them
1. Bare Metal
2. Virtual Machine
3. Docker Compose
4. Kubernetes
5. Managed PAAS (App Engine, Heroku etc)

## Containers

Containers are a way to package our application with all its dependencies. And then we can run it anywhere there is a container runtime available. Docker is one such runtime.

> In GKE the default runtime is Containerd. But your local docker image works just fine because all container images are OCI compliant

```
git clone origin/containers-01
```

```bash
cd rest-api-container
```

Take a look at the `Dockerfile`. This file is the definition on how we want our application packaged. 

To build the Docker image we can run 

```bash
docker build -t <DOCKER_USERNAME>/<IMAGE_NAME>:<TAG>
```

> We are using dockerhub as our image registry. If you are using something different e.g. gcr, quay etc, you will have to put the full URL for docker to find the repo.

Once the build succeeds we can then push the image upto the registry

```bash
docker push <DOCKER_USERNAME>/<IMAGE_NAME>:<TAG>
```

We can also run the container using `docker run`

## Docker Compose

To run our application we can also use docker compose. 

For this example I am using a an external database. I pre-loaded our data in a cloudsql postgres database on GCP. 

To run the `docker-compose` instance all we have to do it

```
docker-compose up
```

This should spin up our application and `cloud-sql-proxy` to connect to our DB instance.

We can access our application as before.

```bash
curl http://localhost:7999/api/v1/users/1
```

Only change here that our database is running on the cloud. 


