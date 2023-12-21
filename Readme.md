# Text To Speech Service

This project provides a text-to-speech conversion service using Kubernetes.

## Deploying the Service
Before deploying the service, you need to build the Docker image and push it to the Google Container Registry (GCR). You can do this by running the following commands:

```bash
make build
make push
```

To deploy the service to a Kubernetes cluster, apply the Kubernetes configuration file included in the project:

```bash
kubectl apply -f kubernetes.yaml
```
### Testing the service
To test the server running on Kubernetes, first get the external IP of the LoadBalancer:

```bash
kubectl get services say-service
```

Next, run the say.go script to test the service. Replace [EXTERNAL-IP] with the external IP address of the service:bash

```
go run cmd/test-user/say.go -b [EXTERNAL-IP]:8080 "Text to test"
```

Next, run the say.go script to test the service. Replace [EXTERNAL-IP] with the external IP address of the service:

```bash
go run cmd/test-user/say.go -b [EXTERNAL-IP]:8080 "Text to test"
```

The output file is saved to test-result/output.wav by default. You can listen to the audio using a media player of your choice.
