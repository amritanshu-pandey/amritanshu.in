---
title: Create self signed certificates for Kubernetes using cert-manager
date: 2021-02-14T08:04:23+11:00
draft: false
categories: [‘kubernetes’]
tags: [‘kubernetes’]
---

## Install Cert manager in Kubernetes

Read this for up-to-date instructions: https://cert-manager.io/docs/installation/kubernetes/

```bash
# Kubernetes 1.16+
$ kubectl apply —validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.0.2/cert-manager.yaml
```

## Create a keypair secret

In this step create a new k8s secret that contains the TLS CA cert and key that is used by cert manager to issue new certificates. As a prerequisite, we need a CA certificate and associated key encoded in base64.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: ca-key-pair
  namespace: default
data:
  tls.crt: <tls-key-base64-encoded>
  tls.key: <tls-key-base64-encoded>
```

## Create an issuer

Issuers are used by Cert manager to issue new certificates

```yaml
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: ca-issuer
  namespace: default
spec:
  ca:
    secretName: ca-key-pair
```

## Create certificates

This creates new certificate using the issuer and CA key pair created earlier. In the following example, the certificate is stored as k8s secret `k8s-xps-lan` in default namespace.

```yaml
apiVersion: cert-manager.io/v1alpha2
kind: Certificate
metadata:
  name: k8s-xps-lan
  namespace: default
spec:
  secretName: k8s-xps-lan
  issuerRef:
    name: ca-issuer
    # We can reference ClusterIssuers by changing the kind here.
    # The default value is Issuer (i.e. a locally namespaced Issuer)
    kind: Issuer
  commonName: k8s.xps.lan
  organization:
  - XPS.LAN
  dnsNames:
  - gitlab.xps.lan
  - minio.xps.lan
  - registry.xps.lan
  - k8s.xps.lan
  - kibana.xps.lan
  - elastic.xps.lan
```

In a separate post, we will see how this certificate can be used by ingress-nginx and other applications.