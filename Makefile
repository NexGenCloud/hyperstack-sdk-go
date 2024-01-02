# Makefile

# Variable for the package name
PACKAGE=gosdk
TIME_PACKAGE=time

# Variable for the OpenAPI file
SPEC=api.json
PROJECT=github.com\/nexgen

# go get subcommands
GET_DEPS=  github.com/deepmap/oapi-codegen/cmd/oapi-codegen


install: # Install all the necessary dependencies
	go get ${GET_DEPS}
	go install ${GET_DEPS}

pull:  # Pulls the latest api.json from the server
	@echo "Pulling the latest swagger spec and generating Go SDK..."
	curl "https://infrahub-api-doc.nexgencloud.com/${SPEC}" | python3 -c "import sys, urllib.parse; print(urllib.parse.unquote(sys.stdin.read()))" > "${SPEC}"

generate: # Generates go sdk file from api.json
	#old & boring :
	#oapi-codegen -package ${PACKAGE} -generate "types,client" ${SPEC} > ${PACKAGE}/${PACKAGE}.go
	#new & stylish
	@echo "Generating Go SDK..."
	if [ -d "./api" ]; then rm -Rf ./api; fi
	if [ -d "./${PACKAGE}" ]; then rm -Rf ./${PACKAGE}; fi
	if [ -d "./${TIME_PACKAGE}" ]; then rm -Rf ./${TIME_PACKAGE}; fi
	go run sdk_generator.go
	mkdir ${PACKAGE}/${TIME_PACKAGE}
	find ./${PACKAGE} -name "*.go" -exec echo sed -i -e "s/*time.Time/*${TIME_PACKAGE}.CustomTime/g" {} \;
	find ./${PACKAGE} -name "*.go" -exec sed -i -e 's/"time"/"${PROJECT}\/${PACKAGE}\/${PACKAGE}\/${TIME_PACKAGE}"/g' {} \;
	sed 's/package-replace-location/'$(TIME_PACKAGE)'/g' CustomTime.stub > ${PACKAGE}/${TIME_PACKAGE}/CustomTime.go

test:
	go test -v sdk_generator_test.go

all: pull install generate test

# The `.PHONY` rule keeps Make from doing something with a file named clean
.PHONY: install pull generate test