NAME            = vicus
DOCKER_IMAGE    = project-x/$(NAME)
# TODO: add version reference to docker image

all: build

.build: .
	docker build -t $(NAME) .
	docker inspect -f '{{.Id}}' $(NAME) > .build

build: .build

release/$(NAME): build
	docker run --rm --entrypoint /bin/sh $(NAME) -c 'tar cf - /$(NAME) /etc/ssl' > $@ || (rm -f $@; false)
	docker build --rm -t $(DOCKER_IMAGE) release

release: release/$(NAME)

push: release
	gcloud docker push $(DOCKER_IMAGE)

clean:
	$(RM) .build release/$(NAME)

proto:
	for f in proto/**/*.proto; do \
		protoc --go_out=plugins=grpc:. $$f; \
		echo compiled: $$f; \
	done

.PHONY: push release build all clean proto
