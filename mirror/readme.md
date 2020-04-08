## Istio Traffic Mirroring

This demo shows how Istio traffic management features can be used to mirror traffic for service without a side car. Mirroring will only occur when traffic source is from a pod that contains the istio sidecar.

Requirements:

1. [Install Istio](../readme.md#Default-Istio-Install)
2. Apply the sample service deployments and only inject the Istio sidecar into Service-A.
	```
	kubectl apply -f namespace.yml -f service-b.yml -f service-c.yml &&  istioctl kube-inject -f service-a.yml| kubectl apply -f -
	```
3. Tail the logs for service-b and service-c deployments.
	```
	$ kubectl logs -f -l app=service-b
	$ kubectl logs -f -l app=service-c
	```
4. Connect your terminal to service-a, install curl and submit a request to service-b. You will see in the logs your request to service-b was mirrored to service-c but your response was returned from service-b only.
	```
	kubectl exec deploy/service-a -- bash -c "apt-get update && apt-get install -y curl; curl service-b:8080"
	```
