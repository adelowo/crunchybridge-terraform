default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

build-dev:
	go install .
	mv ~/go/bin/crunchybridge-terraform ~/go/bin/terraform-provider-crunchybridge
