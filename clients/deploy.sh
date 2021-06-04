./build.sh

docker push minh2804/clients

ssh tom@3.136.26.146 << EOF
	docker stop clients || true && docker rm clients || true

	docker pull minh2804/clients

	docker run -d \
		--name clients \
		-p 80:80 \
		-p 443:443 \
		-v /etc/letsencrypt:/etc/letsencrypt:ro \
		-e TLSCERT=/etc/letsencrypt/live/thenightbeforeitsdue.de/fullchain.pem \
		-e TLSKEY=/etc/letsencrypt/live/thenightbeforeitsdue.de/privkey.pem \
		minh2804/clients

	exit
EOF
