# webtop-k8s-operator

An operator for operating webtop or any desktop/app container in kubernetes.

## Description

This is an operator for [lsio webtop 2.0](https://docs.linuxserver.io/images/docker-webtop/). It is designed to manage desktop container or container app resources.

It also controls the access to the sessions to the user owning the resource using Envoy proxy, OPA and OIDC.

## Scope

Provide basic functionality of creating, updating and deleting resources. Also include persistent storage in some fashion, possibly PVC.
The resources are [lsio webtop](https://docs.linuxserver.io/images/docker-webtop/) containers with [KasmVNC](https://github.com/kasmtech/KasmVNC) on.

## Architecture

### Gateway

The gateway is a backend proxy service made up of Envoy Proxy and Open Policy Agent (OPA).
The gateway can be responsible for multiple sessions for multiple users at the same time. It is also capable of securing the sessions to the owing user based on which namespace the resource was created in and which user owns that namespace.

It utilizes OIDC for authentication and to provide groups and roles. It utilizes OPA for authorization.

Envoy is integrated with OPA to authorize the user and also to control the session. Meaning the gateway is continuously checking the JWT of the user and is capable of cutting the connection if said user is not authorized anymore.

### Desktop

This is a simple linux desktop created by linuxserver.io called webtop. It launches a KasmVNC server inside the container and exposes it on HTTP port 3000 and HTTPS port 3001.

### Application

This is a simple X11 or wayland application running inside a container. It launches a KasmVNC server inside the container and exposes it on HTTP port 3000 and HTTPS port 3001.
An application could for example be firefox.

## Features

- [ ] Create, update, delete webtop instances
- [ ] Persistent user data
- [ ] Proot-apps applications
- [ ] Sleep/pause the pod with inactive use
- [ ] Namespaced resources
- [ ] Secure and private connections

## Getting Started

### Prerequisites

- go version v1.20.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster

**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/webtop-k8s-operator:tag
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/webtop-k8s-operator:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall

**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## License

Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
