# SSH Monitor

## Prequesite
- docker
- footloose
- ansible

## How to run
build and deploy
```bash
ansible-playbook ci/playbook.yaml
```

view ssh attempts
```bash
curl localhost:9999/view
```
