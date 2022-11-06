generate:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/swagger.yaml \
    -g go \
    -o /local/havaclient \
    -c /local/genconfig.yml