docker build -t xuwensi0515/todolistclient .
docker push xuwensi0515/todolistclient
ssh wenxi@3.136.26.146
docker rm -f todolistclient
docker pull xuwensi0515/todolistclient
docker run -d -p 443:443 -p 80:80 -v /etc/letsencrypt:/etc/letsencrypt:ro --name auth xuwensi0515/todolistclient
