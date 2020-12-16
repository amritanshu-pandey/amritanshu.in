---
title: "How to trigger Gitlab pipelines using REST API"
date: 2020-12-16T21:24:00+11:00
draft: true
categories: ['gitlab']
tags: ['gitlab']
---

Gitlab pipelines are the CI(Continuous Integration) build job(s), that are usually triggered when a new commit is pushed to the git repository. Pipelines can be used to test, compile and deploy the code or to automate certain tasks. In addition to changes in git repositories, pipelines can also be triggered through other mechanisms as well. For instance:

1. Manually through the Gitlab UI
1. Based on a schedule
1. Using gitlab's REST API

In this article, I will explain how to trigger a Gitlab pipeline using the Gitlab's REST API.

## Create pipeline trigger
1. Navigate to project settings -> CI/CD -> Pipeline triggers
1. Create a new trigger and make note of the token. Token is used to authenticate the REST calls to trigger the pipeline.
![Gitlab Pipeline Triggers](/static/static/gitlab-pipeline-triggers.png)
1. Also make note of the URL displayed in the same page that must be used to trigger the pipeline. For e.g. `https://git.xps.lan/api/v4/projects/14/trigger/pipeline`

## Trigger pipeline using `curl`

This curl command can be used to trigger the pipeline for project with id `14` for branch or tag specified as `ref` in the curl command

```bash
curl -X POST \
     -F token=<token-noted-in-previous-step> \
     -F "ref=<ref-name-like-branch-or-tag>" \
     https://git.xps.lan/api/v4/projects/14/trigger/pipeline
```

Sometime it might be required to pass on some variables to the pipeline while triggering the same. THe variables passed while triggering the pipeline are avilable as environment variables at run time and takes precedence over all other variables.

Following command will provide a variable names `RUN_NIGHTLY_BUILD` with value as `true` to the pipeline. Any number of variables can be provided like this.

```bash
curl -X POST \
     -F token=<token-noted-in-previous-step> \
     -F "ref=<ref-name-like-branch-or-tag>" \
     -F "variables[RUN_NIGHTLY_BUILD]=true" \
     https://git.xps.lan/api/v4/projects/14/trigger/pipeline
```

## Using Powershell Invoke-WebRequest 

For the users on Windows, it might be bit easier to use Powershell and Invoke-WebRequest, rather than using `curl` on Windows or WSL.

```powershell
Invoke-WebRequest https://git.xps.lan/api/v4/projects/14/trigger/pipeline `
    -Method POST `
    -Body @{"token"="<token-noted-in-previous-step>"; "ref"="<ref-name-like-branch-or-tag>"; "variables[RUN_NIGHTLY_BUILD]"="true"}
```
