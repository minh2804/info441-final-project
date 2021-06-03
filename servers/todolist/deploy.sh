./build.sh

docker push minh2804/todolist

ssh tom@18.117.102.235 << EOF
	docker network create api

	docker rm -f rd
	docker run -d \
		--name rd \
		--network api \
		redis

	docker rm -f todolist
	docker pull minh2804/todolist
	docker run -d \
		--name todolist \
		--network api \
		-p 443:443 \
		-v /etc/letsencrypt:/etc/letsencrypt:ro \
		-e TLSCERT=/etc/letsencrypt/live/api.thenightbeforeitsdue.de/fullchain.pem \
		-e TLSKEY=/etc/letsencrypt/live/api.thenightbeforeitsdue.de/privkey.pem \
		-e SESSIONKEY=qwerty123 \
		-e REDISADDR=rd:6379 \
		-e DSN="root:@tcp(db:3306)/store?parseTime=true" \
		minh2804/todolist

	exit
EOF
