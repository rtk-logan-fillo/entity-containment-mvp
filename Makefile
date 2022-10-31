GITHUB_REPO=entity-containment-mvp
GITHUB_USER=rtk-logan-fillo

install:
	go mod tidy 

validate:
	npx @openapitools/openapi-generator-cli validate -i ./schema/entity-containment.yaml

generate_client:
	npx @openapitools/openapi-generator-cli generate \
		-i ./schema/entity-containment.yaml \
		-g go \
		-o ./entity-containment\
		--git-host github.com \
		--git-repo-id ${GITHUB_REPO} \
		--git-user-id ${GITHUB_USER} \
		--additional-properties=packageName=entitycontainment,isGoSubmodule=true

generate_server:
	npx @openapitools/openapi-generator-cli generate \
		-i ./schema/entity-containment.yaml \
		-g go-server \
		-o .\
		--git-host github.com \
		--git-repo-id ${GITHUB_REPO} \
		--git-user-id ${GITHUB_USER} \
		--additional-properties=packageName=entitycontainment,outputAsLibrary=false,sourceFolder=entity-containment
	for f in entity-containment/*.go; do goimports -w $f; done
