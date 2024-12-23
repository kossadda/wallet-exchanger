.PHONY: mockgen
mockgen:
	mockgen -source=internal/storage/auth/storage.go \
	-destination=internal/storage/auth/mockauth/mockauth.go \
	-package=mockauth
	mockgen -source=internal/storage/wallet/storage.go \
	-destination=internal/storage/wallet/mockwallet/mockwallet.go \
	-package=mockwallet
	mockgen -source=internal/service/grpcclient/service.go \
	-destination=internal/service/grpcclient/mockclient/mockclient.go \
	-package=mockclient