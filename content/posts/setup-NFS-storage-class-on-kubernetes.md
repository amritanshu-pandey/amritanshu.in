---
title: "Setup NFS Storage Class on Kubernetes"
date: 2021-05-29T18:26:51+11:00
draft: false
categories: ['kubernetes']
tags: ['kubernetes', 'nfs']
---

Since the code required for creating various K8S resources for NFS provisioner is split into
multiple files, I prefer to just clone this gist and using `kubectl` to create all resources
in one go 

```bash
$ git clone https://gist.github.com/amritanshu-pandey/8ab00179c98720cbc28d8bb0c7064426
$ cd 8ab00179c98720cbc28d8bb0c7064426
$ kubectl apply -f *.yaml
```

{{< gist amritanshu-pandey 8ab00179c98720cbc28d8bb0c7064426 "deployment.yaml" >}}

{{< gist amritanshu-pandey 8ab00179c98720cbc28d8bb0c7064426 "rbac.yaml" >}}

{{< gist amritanshu-pandey 8ab00179c98720cbc28d8bb0c7064426 "storage-class.yaml" >}}

---


