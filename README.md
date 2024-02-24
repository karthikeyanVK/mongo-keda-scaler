# KEDA example

This repository consists of everything you need to setup simple Kubernetes 
cluster and demonstrate usage of KEDA mongo scalers. For more
samples check https://github.com/kedacore/samples

The included `helper` provides an easy way to perform both 0 -> n and n -> 0 scalings.  

## Create cluster
The deployment consists of 4 components:
- Mongo instance
- Dummy pod that will be scaled up and down
- App service that provides some helper methods
```sh
kubectl apply -f deployment/
```

## Install KEDA
Follow the official KEDA guide https://keda.sh/deploy/


## Observe
To observe how everything works you can watch two things:
- number of pods and their state: `watch -n2 "kubectl get pods"`
- HPA stats: `watch -n2 "kubectl get hpa"`



## Mongo example
To scale the dummy deployment using 
[Mongo scaler]([https://keda.sh/scalers/mysql/](https://keda.sh/docs/2.13/scalers/mongodb/)) first we have to
deploy the `ScaledObjects`:
```sh
kubectl apply -f keda/mongo-hpa.yaml
```
this should result again in creation of `ScaleObject` and an HPA:
```sh
# kubectl get scaledobjects
NAME                 SCALETARGETKIND      SCALETARGETNAME   MIN   MAX   TRIGGERS   AUTHENTICATION          READY   ACTIVE   FALLBACK   PAUSED    AGE
mongo-scaledobject   apps/v1.Deployment   dummy-mongo             5     mongodb    mongodb-local-trigger   True    False    False      Unknown   1d
```

To scale up we have to insert some values to Mongo database. 
To do this we can use the helper app:
```shell script
kubectl exec $(kubectl get pods | grep "server" | cut -f 1 -d " ") -- keda-talk mongo insert
```
to scale down:
```shell script
kubectl exec $(kubectl get pods | grep "server" | cut -f 1 -d " ") -- keda-talk mongo delete
```
and to debug scaling of the dummy pod:
```shell script
kubectl logs -n keda -l app=keda-operator
```
