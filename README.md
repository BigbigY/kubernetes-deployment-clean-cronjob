# Kubernetes-Deployment-Clean-CronJob

[![CircleCI](https://circleci.com/gh/BigbigY/kubernetes-deployment-clean-cronjob.svg?style=shield)](https://circleci.com/gh/BigbigY/kubernetes-deployment-clean-cronjob)

[English](https://github.com/BigbigY/kubernetes-deployment-clean-cronjob/blob/master/README_EN.md) | [中文](https://github.com/BigbigY/kubernetes-deployment-clean-cronjob/blob/master/README.md)

Kubernetes-Deployment-Clean-CronJob 结合Prometheus API 定时清理废弃deployment的任务计划。


## Install

### 你需要安装`glide`来处理依赖

The easiest way to install the latest release on Mac or Linux is with the following script:
```
curl https://glide.sh/get | sh
```
On Mac OS X you can also install the latest release via [Homebrew](https://github.com/Homebrew/homebrew):
```
$ brew install glide
```
On Ubuntu Precise (12.04), Trusty (14.04), Wily (15.10) or Xenial (16.04) you can install from our PPA:
```
sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update
sudo apt-get install glide
```

### 构建`Kubernetes-Deployment-Clean-CronJob`

1.安装依赖
```
make dep
```
2.构建二进制文件
```
make build
```
你也可以直接`make all`, 它包含了`dep`和`build`

构建`docker image`, 在此之前你需要安装依赖
```
make build-docker
```
清除所有构建内容
```
make clean
```

## 使用(Usage)
```
Usage of ./bin/cleanDeployment:
  -web_url string
        (must)HTTP API URL
```
- `web_url`: Prometheus API URL

## 注意

只能在kubernetes集群内部运行, 如果你需要指定kubeconfig的版本, 那么请看exmple中的`cleanDeployment.go`文件