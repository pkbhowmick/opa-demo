# gatekeeper

### Demo

Install gatekeeper using Helm
```bash
$ helm repo add gatekeeper https://open-policy-agent.github.io/gatekeeper/charts
$ helm install gatekeeper/gatekeeper --name-template=gatekeeper --namespace gatekeeper-system --create-namespace
```

Install template & constraints
```bash
$ kubectl apply -f template.yaml
$ kubectl apply -f constraint.yaml
```

Audit the Constraint
```bash
$ kubectl get constraint -A -o yaml
```

Create good & bad ns
```bash
$ kubectl apply -f bad_ns.yaml 
Error from server (Forbidden): error when creating "bad_ns.yaml": admission webhook "validation.gatekeeper.sh" denied the request: [ns-must-have-gk] you must provide labels: {"gatekeeper"}

$ kubectl apply -f good_ns.yaml
namespace/good-ns created
```


