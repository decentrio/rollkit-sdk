###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

# make init-simapp initializes a single local node network
# it is useful for testing and development
# Usage: make install && make init-simapp && simd start
# Warning: make init-simapp will remove all data in simapp home directory

CURRENT_DIR = $(shell pwd)

test-sim-nondeterminism:
	@echo "Running non-determinism test..."
	@cd ${CURRENT_DIR}/simapp && go test -mod=readonly -run TestAppStateDeterminism -Enabled=true \
		-NumBlocks=100 -BlockSize=200 -Commit=true -Period=0 -v -timeout 24h