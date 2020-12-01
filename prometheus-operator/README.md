#### create Minikube cluster
    minikube start --cpus 4 --memory 8192 --vm-driver virtualbox

#### install Prometheus-operator
###### add repos
    helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
    helm repo add stable https://kubernetes-charts.storage.googleapis.com/
    helm repo update

###### install chart
    helm install prometheus prometheus-community/kube-prometheus-stack

###### install chart with fixed version    
    helm install prometheus prometheus-community/kube-prometheus-stack --version "9.4.1"

###### pull up grafana UI locally
kubectl port-forward deployment.apps/prometheus-grafana 3000


