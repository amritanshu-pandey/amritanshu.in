---
title: "Introduction to Prometheus metric types"
date: 2025-01-11T00:00:00+11:00
draft: true
categories: ['prometheus']
tags: ['SRE', 'metrics', 'prometheus']
---

Prometheus is a time series database used for storing metrics and also includes an alerting engine to generate alerts
based on the stored metrics. Prometheus has a rich ecosystem of client libraries and exporters
for either natively producing metrics in Prometheus format or to expose prometheus metrics from applications and services
cannot produce Prometheus format metrics by default.

We will take a detailed look at Prometheus in a separate blog post, and in this post we are going to see in details
the various kind of metrics that Prometheus supports and common scenarios where these metrics types are suitable for.

## Counters

Counter metrics are used for recording observations that grow over time for e.g. number of occurrence of something,
number of requests served, total failures etc. Prometheus only supports cumulative and monotonically increasing counters
i.e. the counter values must always only increase and must never go down in value except for resetting to zero for example
when the process restarts.
