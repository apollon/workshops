.PHONY: clean build buffalo

build: buffalo docker

buffalo:
	cd ./api; GOOS=linux buffalo build -k -o ./app

docker:
	docker build -t shop-api -f ./shop/Dockerfile .
