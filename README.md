## server 启动
```
docker build -f ./Dockerfile-server -t girlfriend-gift-server .
docker run --name="girlfriend-gift-server" --rm -p 8081:8081 -d girlfriend-gift-server 
```


## scanner 运行
```
docker build -f ./Dockerfile-scanner -t girlfriend-gift-scanner .
docker run --rm --name="girlfriend-gift-scanner" -v 本地需要上传图片的绝对路径:/albumDir girlfriend-gift-scanner
```

## schedule_make 运行
```
docker build -f ./Dockerfile-schedule-maker -t girlfriend-gift-schedule-maker .
docker run --rm --name="girlfriend-gift-schedule-maker" girlfriend-gift-schedule-maker
```

## mod_time_update 运行
```
go run app/console/mod_time_update/main.go 本地绝对路径
```