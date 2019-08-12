

构建镜像，传送至目标服务器后load至local repository。
```
docker build . -t rinetd
docekr save rinetd | ssh waret@10.20.1.67 'docker load'
```
