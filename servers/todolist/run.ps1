Start-Process -Path ./build.sh -Wait

docker network create api

docker rm -f app
docker run -d `
	--name app `
	--network api `
	-p 443:443 `
	-v 'C:\Users\Tom:/etc/letsencrypt/live/api.veryoriginaldomain.me:ro' `
	-e TLSCERT=/etc/letsencrypt/live/api.veryoriginaldomain.me/fullchain.pem `
	-e TLSKEY=/etc/letsencrypt/live/api.veryoriginaldomain.me/privkey.pem `
	-e DSN="root:@tcp(db:3306)/store?parseTime=true" `
	minh2804/app
