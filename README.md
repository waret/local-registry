
启动服务侧
```
cd server
docker-compose up -d
```

将代理程序拷贝至端测，并启动守护进程。将本地的5000端口代理至服务侧的端口
```
for i in {21,22,23,24,25}; do scp client-go/revprox root@192.168.2.$i:/root/revprox; done
for i in {21,22,23,24,25}; do ssh root@192.168.2.$i "/root/revprox 192.168.2.10:5000 5000 >/dev/null 2>&1 &"; done
```

```
for i in 1 2 3 4 5 6 7 8; do scp client-go/revprox root@10.20.1.7$i:/root/revprox; done
for i in 1 2 3 4 5 6 7 8; do ssh root@10.20.1.7$i "pkill revprox"; done
for i in 1 2 3 4 5 6 7 8; do ssh root@10.20.1.7$i "/root/revprox 10.20.1.67:5000 5000 >/dev/null 2>&1 &"; done
for i in 1 2 3 4 5 6 7 8; do ssh root@10.20.1.7$i "ps aux | grep revprox"; done
```


```
awk '/dashboard/ {system("sh local-registry/tools/trans-image.sh push "$1" "$2)}' local-registry/lists/images_k8s-v1.15.3.txt
awk '{system("sh local-registry/tools/trans-image.sh push "$1" "$2)}' local-registry/lists/images_k8s-v1.15.3.txt
sh lists/files_k8s-v1.15.3.txt
```


```
sudo yum install -y sshpass
ansible -i 71/inventory.ini all -m ping
ansible-playbook -i kubespray-inventories/71/inventory.ini --become --become-user=root repos/kubespray/cluster.yml
```
