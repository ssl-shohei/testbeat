BUILDID=$(shell git rev-parse HEAD)

.PHONY: all
all: testbeat/deb testbeat/rpm testbeat/darwin testbeat/win testbeat/bin \
	build/upload/build_id.txt

ES_BEATS=../../vendor/github.com/elastic/beats
include $(ES_BEATS)/dev-tools/packer/scripts/Makefile


.PHONY: testbeat
testbeat: build/upload
	# cross compile on ubuntu
	docker run --rm \
		-v $(abspath build):/build \
		-v $(abspath $(ES_BEATS)/dev-tools/packer/xgo-scripts):/scripts \
		-v $(abspath ../..):/source \
		-e PACK=$@ \
		-e BEFORE_BUILD=before_build.sh \
		-e SOURCE=/source \
		-e BUILDID=${BUILDID} \
		tudorg/beats-builder \
		github.com/ssl-shohei//testbeat
	# linux builds on debian 6
	docker run --rm \
		-v $(abspath build):/build \
		-v $(abspath $(ES_BEATS)/dev-tools/packer/xgo-scripts):/scripts \
		-v $(abspath ../..):/source \
		-e PACK=$@ \
		-e BEFORE_BUILD=before_build.sh \
		-e SOURCE=/source \
		-e BUILDID=${BUILDID} \
		tudorg/beats-builder-deb6 \
		github.com/ssl-shohei//testbeat
