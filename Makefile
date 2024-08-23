IMAGE_TAG=danajp/kubernetes-client-example

docker-build:
	docker build -t $(IMAGE_TAG) .

docker-push: docker-build
	docker push $(IMAGE_TAG)
