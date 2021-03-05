# Test 
Test shell scripts 사용법

### jq 설치
* shell 에서 json parsing 시 `jq` 유틸리티를 활용합니다.
* https://stedolan.github.io/jq/

```
$ brew install jq           # mac os
$ sudo apt-get install jq   # linux
```

### CB-Tumblebug 실행

```
$ docker run -d -p 1323:1323 --name cb-tumblebug cloudbaristaorg/cb-tumblebug:v0.x.0-yyyymmdd
```
* 컨테이너 이미지의 최신 tag는 다음을 참조
  * https://hub.docker.com/r/cloudbaristaorg/cb-tumblebug/tags

* 예
```
$ docker run -d -p 1323:1323 --name cb-tumblebug cloudbaristaorg/cb-tumblebug:v0.3.0-espresso
```



## Test 

### CB-Myfw 실행

```
$ cd conf/
$ source setup.env
$ cd ..
$ make
$ make run
```

[Result]
```
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.1.17
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:4321
```

### Create a namespace on CB-Tumblebug
```
$ cd test/2.cluster/
$ ./create-ns.sh
```


### 클러스터 생성
```
$ cd test/1.namespace/
$ ./create-cluster.sh
```

[Success]
```JSON
{
  "name": "myfw-test-cluster",
  "kind": "Cluster",
  "status": "completed",
  "uid": "a66463cb-d027-4376-89e6-d50a82cb6dcb",
  "mcis": "myfw-test-cluster",
  "namespace": "myfw-test-ns",
  "clusterConfig": "",
  "nodes": [
    {
      "name": "myfw-test-cluster-m-lwgc5",
      "kind": "Node",
      "credential": "cb-myfw-test-credential",
      "publicIp": "",
      "uid": "66d47f1e-9220-4151-bead-b1c63c16cdfa",
      "role": "cb-myfw-test-role",
      "spec": "cb-myfw-test-spec"
    }
  ]
}
```

[Already exists]
```JSON
{
  "message": "The cluster myfw-test-cluster already exists."
}
```


### 클러스터 삭제
```
$ ./delete-cluster.sh
```

[Success]
```JSON
{
  "kind": "Status",
  "code": 0,
  "message": ""
}
```

[Does not exist]
```JSON
{
  "message": "The cluster myfw-test-cluster does not exist."
}
```


### 노드 add
```
$ cd test/3.node/
$ ./add-node.sh
```

[Success]
```JSON
{
  "kind": "NodeList",
  "items": [
    {
      "name": "myfw-test-cluster-w-gmwx1",
      "kind": "Node",
      "credential": "cb-myfw-test-credential",
      "publicIp": "",
      "uid": "",
      "role": "cb-myfw-test-role",
      "spec": "cb-myfw-test-spec"
    }
  ]
}
```


### 노드 삭제

```
$ ./delete-node.sh AWS <cluster name> <node name>
```

* 예
```
$ ./delete-node.sh AWS cb-cluster cb-aws-cluster-test-1-w-iqp7n  # AWS
```

[Success]
```JSON
{
  "kind": "Status",
  "code": 1,
  "message": "success"
}
```

[Does not exist]
```JSON
{
  "message": "The node myfw-test-cluster-w-gmwx1 does not exist."
}
```
