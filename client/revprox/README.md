用于将远端HTTP服务映射为本地端口。


构建镜像：
```
docker build . -t revprox
```

通过CMD方式启动转发服务：
```
docker run -it -d -p 5000:5000 revprox www.httpbin.org 5000
```

通过环境变量启动转发服务：
```
docker run -it -d -p 5000:5000 -e 'REMOTE_SVC=www.httpbin.org' -e 'LOCAL_PORT=5000' revprox
```
