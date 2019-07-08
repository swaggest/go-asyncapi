gen:
	json-cli gen-go resources/schema/asyncapi.json --package-name spec --root-name AsyncAPI > spec/entities.go
