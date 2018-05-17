PROFILE:=testaccount
BUCKET:=xcvjalksdfjasieswebsite
REGION:=us-east-1

all: deploy

create:
	@ aws s3 mb "s3://${BUCKET}" --profile "${PROFILE}" --region ${REGION}

# temp flow: yarn build && make deploy
deploy:
	@ aws s3 website "s3://${BUCKET}" --index-document index.html --error-document index.html --region ${REGION} --profile "${PROFILE}"
	@ aws s3 sync build "s3://${BUCKET}" --profile "${PROFILE}" --region ${REGION} --grants read=uri=http://acs.amazonaws.com/groups/global/AllUsers

.PHONY: all deploy