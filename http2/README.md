## Istio Operator

Quickly setup Istio and a simple http/2 service using the Istio operator. This example allows for you to play around with trailers and http2.


#### Create a Kind Cluster

Create a multi-node kind cluster or deploy directly into a managed Kubernetes solution.

Kind cluster configuration `config.yaml`
```
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
- role: worker
- role: worker
```

Kind cluster creation command
```
kind create cluster --config config.yaml --image=kindest/node:v1.17.17
```

#### Istio Operator Install

Install the `istioctl` CLI on your system. You can install using brew `brew install istioctl` or directly from [Istio](https://istio.io/latest/docs/setup/getting-started/#download).

Install the Istio operator but make sure your kube context is pointed at the correct cluster.

```
istioctl operator init
```

This operation will deploy the operator into your local kind cluster under namespace `istio-operator`.

#### Install Istio Mesh

Once the operator is up and running you can install the mesh using `istioctl`.

```
istioctl install --set profile=default --set meshConfig.accessLogFile=/dev/stdout
```

This install command will use the `default` [profile](https://preliminary.istio.io/latest/docs/setup/additional-setup/config-profiles/) which provides `istiod` and `istio-ingressgateway`. This profile is great for getting up and running locally. It's extremely lightweight and only adds a few resources to your cluster.

We added the additional `meshConfig` setting to enable `istio-proxy` sidecar access logs. This makes it really easy to understand what is going on an why your request could be failing.

#### Gateway

Istio should be up and running and ready for use. You can run `istio ps` to check the status of your mesh.

You need to create a simple `Gateway` to allow traffic into the mesh. Apply the gateway resource to the istio-system namespace.

```
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: trailers-gateway
  namespace: istio-system
spec:
  selector:
    istio: ingressgateway # use Istio default gateway implementation
  servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    hosts:
    - "*"
```
*This file is located here [./manifests/istio-gateway.yml](./manifests/istio-gateway.yml)*

#### Startup Services

Now you can begin deploying your services into the cluster. 

Under the [manifests](./manifests) directory you will find a namespace, service, deployment and virtual service. These resources can be deployed directly into your kind cluster. Once the `server` pod is up and running you should be able to port-forward to your `istio-ingressgateway` service using port 80 and submit requests.

Port-Forward to the Istio ingress
```
kubectl port-forward svc/istio-ingressgateway 8080:80
```

Send a pile of requests to the server.
```
$ curl -v localhost:8080/

*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< trailer: Timing
< content-type: text/plain; charset=utf-8
< content-length: 13
< date: Wed, 20 Oct 2021 02:13:44 GMT
< x-envoy-upstream-service-time: 2
< server: istio-envoy
<
* Connection #0 to host localhost left intact
Hello, world!* Closing connection 0
```

#### Additional Info

* Access logs should be enaled if you appled the default profile and meshConfig from above. This really makes troubleshooting a bad configuration a lot easier.
* You can easily review your istio-ingressgateway routes by running `istioctl pc route istio-ingressgateway-7fd7d66b-9r487.istio-system`, of course using your ingres gateway name from `istioctl ps`.
* The docker container contains both `server` and `client` binaries. You can run the client locally pointing at your istio-ingress `./client --addr http://localhost:8080`.


THANK YOU [@stuartcarnie](https://gist.github.com/stuartcarnie) for writing the go code for us to test trailers within Istio!