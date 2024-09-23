ENV              	   ?= stg
CREDENTIAL_FILE_NAME ?= encrypted_secret
VAR_FILE              = ./tfvars/${ENV}.tfvars
VAR_OPTS              = -var-file "$(VAR_FILE)"
BACKEND_FILE 					= ./tfbackend/${ENV}.s3.tfbackend
BACKEND_OPTS          = -backend-config="$(BACKEND_FILE)"

.PHONY: clean
clean:
	rm -rf .terraform

.PHONY: fmt
fmt:
	terraform fmt

.PHONY: init
init:
	terraform init -reconfigure $(BACKEND_OPTS)

.PHONY: plan
plan: init
	terraform plan $(VAR_OPTS) -lock=false -refresh=true

.PHONY: apply
apply: init
	terraform apply $(VAR_OPTS) -lock=false -refresh=true

.PHONY: destroy
destroy: init
	terraform destroy $(VAR_OPTS) -lock=false -refresh=true

.PHONY: show
show:
	terraform show

.PHONY: lint
lint:
	@terraform fmt -check
	@terraform validate
	@tflint --init
	@tflint --recursive

.PHONY: tfstate-bucket
tfstate-bucket:
	aws s3api create-bucket \
		--bucket point-app-tfstate-stg \
		--create-bucket-configuration LocationConstraint=ap-southeast-1 \
		--endpoint-url=http://127.0.0.1:5000
	aws s3api put-bucket-versioning \
		--bucket point-app-tfstate-stg \
		--versioning-configuration Status=Enabled \
		--endpoint-url=http://127.0.0.1:5000
