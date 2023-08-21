#!/usr/bin/env bash

find ./bindings/ -name .git -prune -o -type f -not -name prepare_bindings.sh -print0 | xargs -0 perl -i -pe "s|https://raw.githubusercontent.com/asyncapi/asyncapi-node/v2.7.7/schemas/2.0.0.json#/definitions/schema|http://asyncapi.com/definitions/2.5.0/schema.json|g"
find ./bindings/ -name .git -prune -o -type f -not -name prepare_bindings.sh -print0 | xargs -0 perl -i -pe "s|https://raw.githubusercontent.com/asyncapi/asyncapi-node/v2.7.7/schemas/2.0.0.json#/definitions/specificationExtension|http://asyncapi.com/definitions/2.5.0/specificationExtension.json|g"
find ./bindings/ -name .git -prune -o -type f -not -name prepare_bindings.sh -print0 | xargs -0 perl -i -pe 's|"title"|"x-title"|g'



