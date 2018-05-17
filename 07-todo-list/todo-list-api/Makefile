APPLICATION_NAME := todo-list
         LAMBDAS := delete-todo get-todo post-todo put-todo
      AWSPROFILE := testaccount
     	S3BUCKET := xcvjalksdfjasies
	   AWSREGION := us-east-1

.PHONY: deps clean build package deploy local-start-api

deps:
	go get -u ./...

clean: 
	rm -rf ./hello-world/hello-world
	
build:
	for l in $(LAMBDAS); \
	do \
		echo building $$l $$(date -u); \
		GOOS=linux GOARCH=amd64 go build -o $$l src/cmd/$$l/$$l.go; \
		echo completed building $$l $$(date -u); \
		ls -lh $$l-handler; echo \\n; \
	done
	@zip "${APPLICATION_NAME}.zip" *-handler;

package: build
	samdev package --template-file template.yaml --s3-bucket ${S3BUCKET} --profile ${AWSPROFILE} > build-output.yaml

deploy: package
	sam deploy --template-file build-output.yaml --stack-name ${APPLICATION_NAME} --capabilities CAPABILITY_IAM --profile ${AWSPROFILE} --region ${AWSREGION}

local-start-api: build
	samdev local start-api -t template.yaml 
