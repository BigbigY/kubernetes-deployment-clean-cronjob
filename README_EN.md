# Kubernetes-Deployment-Clean-CronJob

[![CircleCI](https://circleci.com/gh/BigbigY/kubernetes-deployment-clean-cronjob.svg?style=shield)](https://circleci.com/gh/BigbigY/kubernetes-deployment-clean-cronjob)

[English](https://github.com/BigbigY/kubernetes-deployment-clean-cronjob/blob/master/README_EN.md) | [中文](https://github.com/BigbigY/kubernetes-deployment-clean-cronjob/blob/master/README.md)

kubernetes-deployment-clean-cronjob task plan for regularly cleaned waste Deployment.

## Install

### You need to install `glide` to deal with

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

### Build `kubernetes-deployment-clean-cronjob`

1.Install dependencies
```
make dep
```
2.Build binaries
```
make build
```
You can also directly `make all`, it contains a `dep` and `build`

Build `docker image`, before that you need to install the dependency
```
make build-docker
```
Clear all build content
```
make clean
```

## Usage
**probe interval is 3 seconds**
```
Usage of ./bin/cleanDeployment:
  -web_url string
        (must)HTTP API URL
```
- `web_url`: Prometheus API URL

## Please note that

Can only be run inside the kubernetes cluster, if you need to specify the kubeconfig version, so please look at the `exmple/exmple_cleanDeployment.go` file