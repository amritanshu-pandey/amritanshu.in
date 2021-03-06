---
title: "Systemd-Journald"
date: 2020-02-28T03:34:23+11:00
draft: false
categories: ['systemd', 'linux']
tags: ['journald']
---

`journald` is a part of the `systemd` suite of softwares that provide variety of system components for Linux based
operating systems. journald is a daemon responsible for recording logs generated by Linux kernel and other
applications installed on the system.

Since its inception and later mass adoption by majority of the Linux distributions, systemd has remained a
controversial topic among the Linux enthusiasts and sysadmins. Now there are both valid and not-so-sound arguments
against systemd; but the fact that almost every major Linux distribution today uses systemd speaks for its merit
and is a reason good enough for us to understand it better.

I for one quite like the convenience and coherent ecosystem of utilities and commands that systemd provides and
in this series of web posts I want to share some of the systemd related tid-bits and the use-cases I solve using
systemd utilities.

This first post is about the log collection and management utility provided by the systemd i.e. `systemd-journald`.


## What is `systemd-journald`

`systemd-journald` is a daemon to collect event logs into its data store in binary format. This daemon can be
configured by modifying the configuration file that by default is stored at `/etc/systemd/journald.conf`.

Configuration files for all systemd daemons (including journald) are in `INI` format and are stored under
`/etc/systemd/`. After modifying the `INI` files it is necessary to restart the corresponding daemon using `systemctl`

#### How to restart journald
```bash
systemctl restart systemd-journald
```

Detailed reference for the journald config can be found at the freedesktop website here - https://www.freedesktop.org/software/systemd/man/journald.conf.html#

Few facts about the journal daemon and its config that stand out for me are:

1. journald can be configured to store the logs in either the persistent file system under the default path of
    `/var/log/journal/` or in a volatile file system under  `/run/log/journal` by default. Logs stored in the
    persistent file system are stored permanently depending upon the retention factors specified in config like total
    size, age etc. Logs stored in the volatile file system also obey similar retention logic, but are removed upon
    reboot.

2. journald can be configured to forward the logs to syslog. This the default behaviour for most Linux distributions.
   For e.g. in fedora journald logs are also stored in syslog format in plain text in various files under `/var/log`.

   ```bash
   $ sudo tail -3 /var/log/messages

   Feb 28 11:49:03 xpsws NetworkManager[773]: <info>  [1582850943.6441]
   Feb 28 11:49:03 xpsws systemd[1]: Starting Network Manager
   Feb 28 11:49:03 xpsws systemd[1]: Started Network Manager
   ``` 

## What is `journalctl`

Journalctl is a command line interface to query the logs stored by journald in its binary store.

When invoked without any parameters, it displays all the logs in default syslog format. Following are a few more
examples of the common journalctl commands that highlight its powerful features.

### Common commands

- To follow all the logs on stdout (analogous to `tail -f /var/log/syslog`)

```bash
$ journalctl -f
```

- To display logs pertaining to a service (systemd unit) and display only last 10 log entries

```bash
$ journalctl -u <servicename> -n 10
```

- To display logs in a different format e.g. `json`

```bash
$ journalctl -u <servicename> -n 2 -o json
```

 Notable supported formats: json, json-pretty, short, short-iso, verbose

- To display the amount of disk used by journald logs

```bash
$ journalctl --disk-usage

Archived and active journals take up 24.0M in the file system.
```

A detailed reference guide for the journalctl cli can be found here - https://www.freedesktop.org/software/systemd/man/journalctl.html