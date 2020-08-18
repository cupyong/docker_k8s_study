Docker的应用场景

> 1.简化配置,同一Docker的配置可以在不同 环境中使用,降低了硬件要求和应用环境之间的耦合度. 
> 2.代码的流水线管理.代码从开发者的假期到最终在生产环境上的部署,需要经过很多的中间环境.而每一个中间环境都有自己微小的蛤贝,Docker给应用提供一个从开发到上线均一致的环境,让代码的流水线变得简单. 
> 3.提高开发效率 
> 4.隔离应用,使应用松耦合 
> 5.快速部署 
> docker [CMD][options] 
> 基本命令与释义 
> attach进入到正在运行的容器 
> build由Dockerfile构建镜像 
> commit由容器的改变创建一个新的镜像 
> cp在容器中复制文件或文件夹到本地文件或文件夹中. 
> logs获取容器日志 
> network管理Docker网络 
> node管理Docker集群节点 
> pause暂停一个或多个容器内的所有进程 
> port列表端口映射或用于容器的特定的映射 
> ps列出容器 
> pull从镜像仓库中拉出镜像 
> push上传镜像 
> rename重命名镜像 
> restart重启一个容器 
> rm 删除容器 
> rmi删除镜像 
> run在容器中运行命令 
> search在Docker Hub中查找镜像 
> service管理Docker服务 
> start启动停止的容器 
> stats显示容器的实时流资源使用统计信息 
> stop停止正在运行的容器 
> swarm管理Docker集群 
> tag将镜像标记到存储库中 
> top显示容器的正在运行的进程 
> volume管理Docker卷

docker run [options] 
常用参数与释义（主要介绍docker run）

> -a stdin: 指定标准输入输出内容类型，可选 STDIN/STDOUT/STDERR 三项； 
> -d: 后台运行容器，并返回容器ID； 
> -i: 以交互模式运行容器，通常与 -t 同时使用； 
> -t: 为容器重新分配一个伪输入终端，通常与 -i 同时使用； 
> –name=”nginx-lb”: 为容器指定一个名称； 
> –dns 8.8.8.8: 指定容器使用的DNS服务器，默认和宿主一致； 
> –dns-search example.com: 指定容器DNS搜索域名，默认和宿主一致； 
> -h “mars”: 指定容器的hostname； 
> -e username=”ritchie”: 设置环境变量； 
> –env-file=[]: 从指定文件读入环境变量； 
> –cpuset=”0-2” or –cpuset=”0,1,2”: 绑定容器到指定CPU运行； 
> -m :设置容器使用内存最大值； 
> –net=”bridge”: 指定容器的网络连接类型，支持 bridge/host/none/Container: 四种类型； 
> –link=[]: 添加链接到另一个容器； 
> –expose=[]: 开放一个端口或一组端口；



进入容器

```
docker run -i -t tomcat /bin/bash
```

下载镜像