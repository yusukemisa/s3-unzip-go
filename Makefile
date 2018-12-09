STACK_NAME := stack-unzipper-lambda
TEMPLATE_FILE := template.yml
SAM_FILE := sam.yml

build:
	GOARCH=amd64 GOOS=linux go build -o artifact/unzipper
.PHONY: build

deploy: build
	sam package --template-file $(TEMPLATE_FILE) --s3-bucket $(STACK_BUCKET) --output-template-file $(SAM_FILE)
	sam deploy --template-file $(SAM_FILE) --stack-name $(STACK_NAME) --capabilities CAPABILITY_IAM --parameter-overrides ZippedArtifactBucket=$(ZIPPED_ARTIFACT_BUCKET) UnzippedArtifactBucket=$(UNZIPPED_ARTIFACT_BUCKET)
.PHONY: deploy

s3-init:
	aws s3 mb "s3://$(STACK_BUCKET)"
.PHONY: s3-init

delete:
	aws s3 rm "s3://$(ZIPPED_ARTIFACT_BUCKET)" --recursive
	aws s3 rm "s3://$(UNZIPPED_ARTIFACT_BUCKET)" --recursive
	aws cloudformation delete-stack --stack-name $(STACK_NAME)
	aws s3 rm "s3://$(STACK_BUCKET)" --recursive
	aws s3 rb "s3://$(STACK_BUCKET)"
.PHONY: delete

zip-clean:
	aws s3 rm "s3://$(ZIPPED_ARTIFACT_BUCKET)/names.zip"
.PHONY: zip-clean

zip-up:
	aws s3 cp "./testdata/names.zip" "s3://$(ZIPPED_ARTIFACT_BUCKET)/names.zip"
.PHONY: zip-up

test:
	go test ./...
.PHONY: test