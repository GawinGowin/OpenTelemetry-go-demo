all:
	export OTEL_RESOURCE_ATTRIBUTES="service.name=dice,service.version=0.1.0" && \
	export OTEL_EXPORTER_OTLP_ENDPOINT="http://collector:4317" && \
	go mod tidy && go run .
