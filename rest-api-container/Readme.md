# Application Delivery

The best application is a running application. For our we worked on creating some amazing REST API so far but its no use to anyone if we don't have it accessible to others.

## Kubernetes

We can run the docker-compose in a virtual machine and expose it onto the internet that way. But these days for many software teams Kubernetes is the defacto place to deploy containers. Kubernetes is an open-source tool for managing container workloads. It has some awesome properties like self healing and auto-scaling. Its also the center of many cloud native technologies that makes the application delivery experience even better.

For this example we will deploy to a GKE Standard cluster. To enable our `cloud-sql-proxy` to connect to our database we will have to set up some IAM policies [as shown here](https://cloud.google.com/sql/docs/postgres/connect-kubernetes-engine#service-account-key)

```bash
git clone origin/containers-02
```

If you are not already in the folder

```bash
cd rest-api-container
```

```bash
kubectl apply -f kubernetes/
```

After a while we should see that our application is deployed

```bash
kubectl get deploy
```

```bash
NAME      READY   UP-TO-DATE   AVAILABLE   AGE
restapi   1/1     1            1           64m
```

We can get the external IP from our svc

```bash
kubectl get svc
```

```bash
NAME         TYPE           CLUSTER-IP    EXTERNAL-IP    PORT(S)        AGE
kubernetes   ClusterIP      10.52.0.1     <none>         443/TCP        5d8h
restapi      LoadBalancer   10.52.7.179   <SOME_IP>	     80:30981/TCP   69m
```

We can access our application at the IP

```bash
curl -s http://<SOME_IP>/api/v1/users/1 | jq
```

We should get

```json
{
  "id": 1,
  "name": "Travis Miller",
  "email": "michellebrooks@williams.net",
  "company": "Russell-Rowe",
  "interests": [
    "Jenkins",
    "Jupyter Notebook",
    "Swift",
    "Scrum",
    "Spark",
    "Kubeflow",
    "GraphQL",
    "CSS",
    "Java",
    "Go",
    "Openshift",
    "C++",
    "Data Encryption"
  ]
}
```

