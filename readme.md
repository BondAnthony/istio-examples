# Istio Examples

Just some examples on how you can use Istio.

[Mirroring traffic](./mirror/readme.md) to services without sidecars.


## Default Istio Install

1. Download Istio, this will download the latest release to your local directory and extract.
	```
	curl -L https://istio.io/downloadIstio | sh -
	```
2. From within the `istio-[version]` directory install Istio `default` profile into a running cluster.
	```
	./bin/istioctl manifest apply --set profile=default
	```