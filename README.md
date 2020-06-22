# AMRITANSHU.IN

My blog at [https://blog.amritanshu.in]

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
