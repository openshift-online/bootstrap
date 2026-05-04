.PHONY: lint clean help

lint:
	shellcheck scripts/*.sh 2>/dev/null || echo "shellcheck not installed"

clean:
	@echo "Nothing to clean"

help:
	@echo "bootstrap - Regional foundation for cloud services at Red Hat"
	@echo ""
	@echo "Targets:"
	@echo "  lint   - Run shellcheck on scripts"
	@echo "  clean  - Clean build artifacts"
	@echo "  help   - Show this help"
