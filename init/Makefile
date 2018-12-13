
# Backing Service Name
BACKING_SERVICE_NAME ?= BackingService

all: create

create:
	cd .. && kubebuilder init --domain sap.com --license apache2 --owner "The K8S Backing Service Authors" --depArgs -v
	cd .. && kubebuilder create api --group backingservices --version v1alpha1 --kind ${BACKING_SERVICE_NAME}
	sed 's/<BACKING_SERVICE>/${BACKING_SERVICE_NAME}/g' service_types.go > ../pkg/apis/backingservices/v1alpha1/$(shell echo $(BACKING_SERVICE_NAME) | tr A-Z a-z)_types.go
	mkdir ../pkg/utils
	sed 's/<BACKING_SERVICE>/${BACKING_SERVICE_NAME}/g' util.go > ../pkg/utils/util.go

