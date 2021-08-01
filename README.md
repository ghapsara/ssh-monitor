# SSH Monitor

## Prequesite
- docker
- footloose
- ansible

### Deploy footloose machine
```bash
footloose create
```

## How to run
build and deploy
```bash
ansible-playbook ci/playbook.yaml
```

view ssh attempts
```bash
curl localhost:9999/view
```
