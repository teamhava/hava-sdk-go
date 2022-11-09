.PHONY: validate

validate:
    docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli validate \
    -i /local/swagger.yaml

generate:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/generatorconf/swagger.yaml \
    -g go \
    -o /local/ \
    -c /local/generatorconf/genconfig.yml \
    -t /local/generatorconf/templates