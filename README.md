# SSH Monitor

## Prequesite
- [docker](https://www.docker.com/)
- [footloose](https://github.com/weaveworks/footloose)
- [ansible](https://www.ansible.com/)

## How to run
build and deploy
```bash
ansible-playbook ci/playbook.yaml
```

view ssh attempts
```bash
curl localhost:9999/view
```
