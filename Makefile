.PHONY: all clean config run test

ifeq ($(OS), Windows_NT)
	CONFIG_SCRIPT = scripts\config.sh
else
	CONFIG_SCRIPT = ./scripts/config.sh
endif

config:
	@$(CONFIG_SCRIPT)
	# FIXME: The above doesn't work?
	./scripts/config.sh

run: config
	@go run main.go $(ARGS)
