
kubectl create secret generic mysql-pass --from-literal=password="${1?password}"

