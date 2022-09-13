API_DIR = cmd

AZURE_API_NAME_PROD = coaty-world-arweave

deploy-prod-arweave:
	@GOOS=linux GOARCH=amd64 go build $(API_DIR)/main.go && mv main $(API_DIR)/functions
	(cp config.yaml $(API_DIR)/functions/config.yaml && cd $(API_DIR)/functions && func azure functionapp publish $(AZURE_API_NAME_PROD))
	@rm -rf $(API_DIR)/functions/main
	@rm -rf $(API_DIR)/functions/config.yaml
	@echo "Deployed prod $(AZURE_API_NAME_PROD)"
