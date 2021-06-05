./build.sh

docker push minh2804/app

ssh tom@3.136.26.146 << EOF
	docker stop app || true && docker rm app || true

	docker pull minh2804/app

	docker run -d \
		--name app \
		-p 80:80 \
		-p 443:443 \
		-v /etc/letsencrypt:/etc/letsencrypt:ro \
		-e TLSCERT=/etc/letsencrypt/live/thenightbeforeitsdue.de/fullchain.pem \
		-e TLSKEY=/etc/letsencrypt/live/thenightbeforeitsdue.de/privkey.pem \
		minh2804/app

	exit
EOF
