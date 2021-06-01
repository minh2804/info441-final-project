./build.sh

docker push minh2804/db

ssh tom@18.117.102.235 << EOF
	docker network create api

	docker rm -f db
	docker pull minh2804/db
	docker run -d \
		--name db \
		--network api \
		-e MYSQL_ALLOW_EMPTY_PASSWORD=true \
		-e MYSQL_DATABASE=store \
		minh2804/db
	exit
EOF
