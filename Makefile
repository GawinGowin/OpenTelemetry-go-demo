
all:
	export OTEL_RESOURCE_ATTRIBUTES="service.name=dice,service.version=0.1.0"
	cd src && go mod tidy && go run .