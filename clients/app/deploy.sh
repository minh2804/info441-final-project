./build.sh

docker push minh2804/client

ssh tom@3.136.26.146 << EOF
	docker stop client || true && docker rm client || true

	docker pull minh2804/client

	docker run -d \
		--name client \
		-p 80:80 \
		-p 443:443 \
		-v /etc/letsencrypt:/etc/letsencrypt:ro \
		-e TLSCERT=/etc/letsencrypt/live/thenightbeforeitsdue.de/fullchain.pem \
		-e TLSKEY=/etc/letsencrypt/live/thenightbeforeitsdue.de/privkey.pem \
		minh2804/client

	exit
EOF
