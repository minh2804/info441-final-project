docker rm -f rd

docker run -d `
	--name rd `
	--network api `
	-p 6379:6379 `
	redis
