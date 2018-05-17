APPLICATION_NAME := todo-list
         LAMBDAS := delete-todo get-todo post-todo put-todo
      AWSPROFILE := testaccount
     	S3BUCKET := xcvjalksdfjasies
	   AWSREGION := us-east-1

.PHONY: deps clean build package validate deploy local-start-api

# temp flow: make package && make deploy
all: deploy clean

deps:
	go get -u ./...

clean: 
	rm -rf ./hello-world/hello-world
	
build:
	for l in $(LAMBDAS); \
	do \
		echo building $$l $$(date -u); \
		GOOS=linux GOARCH=amd64 go build -o $$l-handler src/cmd/$$l/$$l.go; \
		echo completed building $$l $$(date -u); \
		ls -lh $$l-handler; echo \\n; \
	done
	@zip "${APPLICATION_NAME}.zip" *-handler;

package:
	samdev package --template-file template.yaml --s3-bucket ${S3BUCKET} --profile ${AWSPROFILE} > build-output.yaml

validate:
	@bash tests/validate-templates ${AWSPROFILE}

deploy: validate
	sam deploy --template-file build-output.yaml --stack-name ${APPLICATION_NAME} --capabilities CAPABILITY_IAM --profile ${AWSPROFILE} --region ${AWSREGION}

local-start-api: build
	samdev local start-api -t template.yaml 
