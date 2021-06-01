Start-Process -Path ./build.sh -Wait

docker network create api

docker rm -f db
docker run -d `
	--name db `
	--network api `
	-e MYSQL_ALLOW_EMPTY_PASSWORD=true `
	-e MYSQL_DATABASE=store `
	minh2804/db

# docker run -d `
# 	--name db `
# 	-p 3306:3306 `
# 	-e MYSQL_ROOT_PASSWORD=test `
# 	-e MYSQL_DATABASE=store `
# 	minh2804/db

# docker run -it `
# 	--rm `
# 	--network host `
# 	mysql sh -c "mysql -h127.0.0.1 -uroot -ptest"
