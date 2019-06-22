carbonara
=========
run command backup app.

# carbonara config example

## simple ssh
```yaml
hosts:
  - hostname: "gateway"
    addr: "192.168.1.1"
    username: "admin"
    password: "admin"
    ssh: true
    #webhook: slack webhook url
    cmds:
      - "show configuration | no-more"
```

## enable password
```
hosts:
  - hostname: "backbone-router"
    addr: "192.168.99.1"
    username: "admin"
    password: "admin"
    ssh: true
    shell: true
    #webhook: slack webhook url
    shell_wait: 3
    cmds_pre:
      - "enable"
      - "ENABLE-PASSWORD"
    cmds:
      - "show run | no-more"
    cmds_post:
      - "exit"
```
