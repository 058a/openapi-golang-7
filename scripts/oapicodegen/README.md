~/go/bin/oapi-codegen -generate types -package hello ./../../api/hello.yaml > ./../../internal/infra/oapicodegen/hello/hello.go && go mod tidy

~/go/bin/oapi-codegen -generate types -package item ./../../api/stock/item.yaml > ./../../internal/infra/oapicodegen/stock/item/item.go && go mod tidy

~/go/bin/oapi-codegen -generate types -package location ./../../api/stock/location.yaml > ./../../internal/infra/oapicodegen/stock/location/location.go && go mod tidy

~/go/bin/oapi-codegen -generate types -package unit ./../../api/stock/unit.yaml > ./../../internal/infra/oapicodegen/stock/unit/unit.go && go mod tidy