# opa-demo

## Check any policy using OPA Go library 
- Example: ```$ go run main.go --query=data.example.no_slow_query --input=input.json --regofile=example.rego```


## API Authorization using OPA

- ```$ kubectl create configmap example-policy --from-file ./authz/authz.rego```
- ```$ kubectl apply -f deploy/```
- ```$ kubectl port-forward service/opa 8181:8181```
- ```$ cd server```
- ```$ go run main.go```  
- ```$ curl --user alice:demo localhost:8081/salary/bob```  [forbidden]
- ```$ curl --user alice:demo localhsot:8081/salary/alice``` [permitted]


## Resources
- [OPA Official Doc](https://www.openpolicyagent.org/docs/latest/)
- [OPA Github repo](https://github.com/open-policy-agent/opa)  
- [Intro: Open Policy Agent - Torin Sandall, Styra](https://www.youtube.com/watch?v=Lca5u_ODS5s)
- [Open Policy Agent (OPA) Intro & Deep Dive - Anders Eknert, Styra & Will Beason, Google](https://www.youtube.com/watch?v=MhyQxIp1H58)
- [Rego Playground](https://play.openpolicyagent.org/)
- [Example Go Service using OPA for API authz](https://github.com/open-policy-agent/example-api-authz-go)
