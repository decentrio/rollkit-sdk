
###############################################################################
###                                Protobuf                                 ###
###############################################################################

DOCKER := $(shell which docker)
HTTPS_GIT := https://github.com/decentrio/rollkit-sdk.git

containerProtoVer=0.13.0
containerProtoImage=ghcr.io/cosmos/proto-builder:$(containerProtoVer)

protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage)


proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh;

proto-check:
	@if git diff --quiet --exit-code main...HEAD -- proto; then \
		echo "Pass! No committed changes found in /proto directory between the currently checked out branch and main."; \
	else \
		echo "Committed changes found in /proto directory between the currently checked out branch and main."; \
		modified_protos=$$(git diff --name-only main...HEAD proto); \
		modified_pb_files= ; \
        for proto_file in $${modified_protos}; do \
            proto_name=$$(basename "$${proto_file}" .proto); \
            pb_files=$$(find x/ccv -name "$${proto_name}.pb.go"); \
            for pb_file in $${pb_files}; do \
                if git diff --quiet --exit-code main...HEAD -- "$${pb_file}"; then \
                    echo "Missing committed changes in $${pb_file}"; \
					exit 1; \
                else \
                    modified_pb_files+="$${pb_file} "; \
                fi \
            done \
        done; \
		echo "Pass! Correctly modified pb files: "; \
		echo $${modified_pb_files}; \
    fi

proto-format:
	@echo "Formatting Protobuf files"
	@$(protoImage) find ./ -name "*.proto" -exec clang-format -i {} \;

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@$(protoImage) sh ./scripts/protocgen.sh

proto-lint:
	@$(protoImage) buf lint --error-format=json

proto-check-breaking:
	@$(protoImage) buf breaking --against $(HTTPS_GIT)#branch=main

proto-update-deps:
	@echo "Updating Protobuf dependencies"
	$(protoImage) buf mod update

.PHONY: proto-all proto-gen proto-format proto-lint proto-check proto-check-breaking proto-update-deps mocks