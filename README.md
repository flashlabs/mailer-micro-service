# Mailer Micro Service

Micro service that stores customers and send them emails based on `mailing ID`.

## Requirements

- `docker`
- `kubectl`
- `minikube`

## Run the project

Build the app image:

```shell
docker build -t mailer-micro-service-app:latest .
```

Start the Kubernetes cluster:

```shell
make mks
```

Start the application:

```shell
make ks
```

Start the proxy on `127.0.0.1:8080`(*)

```shell
make kbpa
```

To connect manually to the database, open connection on `localhost:5432` (*)

```shell
make kbpp
```

*) Please be sure that the node is in `Running` state. You can check it with the `kubectl get pods` command.

## Test the project 

Create mailing records:

```shell
curl -X POST localhost:8080/api/messages -d '{"email":"john.doe@example.com","title":"Gummi bears","content":"Wafer tart cupcake carrot cake icing","mailing_id":1, "insert_time": "2020-04-24T05:42:38.725412916Z"}'

curl -X POST localhost:8080/api/messages -d '{"email":"john.doe@example.com","title":"Chocolate","content":"Sugar plum gingerbread biscuit gummies cotton candy","mailing_id":2, "insert_time": "2020-04-24T05:42:38.725412916Z"}'

curl -X POST localhost:8080/api/messages -d '{"email":"john.doe@example.com","title":"Marshmallow","content":"Jelly beans candy topping jelly beans sugar plum","mailing_id":3, "insert_time": "2020-04-24T05:42:38.725412916Z"}'
```

Send a mocked message to everyone with given `mailing ID`:

```shell
 curl -X POST localhost:8080/api/messages/send -d '{"mailing_id":1}'
```

Delete a given mailing:

```shell
curl -X DELETE localhost:8080/api/messages/{id}
```

## Shutdown the project

Stop the application:

```shell
make kd
```

Stop the cluster:

```shell
make mkd
```
