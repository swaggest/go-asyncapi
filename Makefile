# json-cli can be found at https://github.com/swaggest/json-cli/releases
gen:
	@json-cli gen-go resources/schema/asyncapi.json --package-name spec --root-name AsyncAPI > spec/entities.go
