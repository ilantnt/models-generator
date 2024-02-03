.PHONY: run

generate:
	go run parser.go
sync:
	tscriptify -package=models_generator/generated_models -target=Frontend/models/models.ts -interface models.go