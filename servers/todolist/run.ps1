Start-Process -Path ./build.sh -Wait

docker network create api

# -p 443:443 `

docker rm -f todolist
docker run -d `
	--name todolist `
	--network api `
	-p 80:80 `
	-v 'C:\Users\Tom:/etc/letsencrypt/live/api.veryoriginaldomain.me:ro' `
	-e TLSCERT=/etc/letsencrypt/live/api.veryoriginaldomain.me/fullchain.pem `
	-e TLSKEY=/etc/letsencrypt/live/api.veryoriginaldomain.me/privkey.pem `
	-e SESSIONKEY=qwerty123 `
	-e REDISADDR=rd:6379 `
	-e DSN="root:@tcp(db:3306)/store?parseTime=true" `
	minh2804/todolist
