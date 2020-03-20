# AMRITANSHU.IN [![Netlify Status](https://api.netlify.com/api/v1/badges/a2854d86-c2f5-4bfd-9c4f-99e771a6430f/deploy-status)](https://app.netlify.com/sites/amritanshu/deploys)

# How to deploy

## Website
Web changes are deployed automatically using the systemd services and timeres.
The timer pulls git changes periodically (30s at time of writing).

## Systemd changes
1. Wait for the git changes to be pulled by systemd timer or pull the changes
   manually
   ```bash
   $ git fetch && git rebase
   ```
1. Copy the `systemd` services and timer units to `/etc/systemd/system/`
2. Reload the `systemd` daemon
```bash
$ systemctl daemon-reload
```
3. Restart `systemd` timers and services
```bash
$ systemctl restart hugo-amritanshu-in.service
$ systemctl restart refresh-hugo-amritanshu-in.timer
```
