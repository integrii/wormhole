default: build push

build:
	docker build -t integrii/wormhole .

push:
	docker push integrii/wormhole
