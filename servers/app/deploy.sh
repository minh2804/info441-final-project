./build.sh

docker push minh2804/app

ssh tom@18.117.102.235 << EOF
	docker network create api

	docker rm -f app
	docker pull minh2804/app
	docker run -d \
		--name app \
		--network api \
		-p 443:443 \
		-v /etc/letsencrypt:/etc/letsencrypt:ro \
		-e TLSCERT=/etc/letsencrypt/live/api.veryoriginaldomain.me/fullchain.pem \
		-e TLSKEY=/etc/letsencrypt/live/api.veryoriginaldomain.me/privkey.pem \
		-e DSN="root:@tcp(db:3306)/store?parseTime=true" \
		minh2804/app

	exit
EOF
