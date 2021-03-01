if [ "$MY_LIB_K8S_DEFINED" = "" ]; then
  export MY_LIB_K8S_DEFINED=1
  source "$(brew --prefix)/opt/kube-ps1/share/kube-ps1.sh"
  # export PS1='$(kube_ps1)'$PS1
fi

switch_k8s_namespace() {
  local ns=${1?namespace}
  kubectl config \
    set-context \
    $(kubectl config current-context) \
    --namespace="$ns"
}

switch_k8s_context() {
  local ctxt=${1?context name}
  kubectl config \
    use-context "${ctxt}"
}

k8s_churn_role() {
  local role=${1?role}
  shift

  local context=${K8S_CONTEXT:-}
  local context_arg=""
  if [ "$context" != "" ]; then
    context_arg="--context=$context"
  fi

  local namespace=${K8S_NAMESPACE:-}
  local namespace_arg=""
  if [ "$namespace" != "" ]; then
    namespace_arg="--namespace=$namespace"
  fi

  local kubectl_cmd="kubectl $context_arg $namespace_arg"

  local all_pods="$($kubectl_cmd get pods | awk "\$1~/^$role-[a-f0-9]+-[a-z0-9]+\$/{printf \" pod/\"\$1;}")"
  echo "$kubectl_cmd delete $all_pods"
  $kubectl_cmd delete $all_pods
}

k8s_search_ip() {
  local ip=${1?ip}
  kubectl get pods --field-selector status.podIP="$ip"  --all-namespaces
}

k8s_ssh() {
  local role="${1?role}"
  local kubectl_cmd="kubectl"

  # local all_pods="$($kubectl_cmd get pods | awk "\$1~/^$role-[a-f0-9]+-[a-z0-9]+\$/{printf \" pod/\"\$1;}")"
  local pod_name="$($kubectl_cmd get pods | awk "\$1~/^$role-[a-f0-9]+-[a-z0-9]+\$/{print \$1}" | head -n 1)"
  echo "pod name: $pod_name"
  $kubectl_cmd exec -ti "${pod_name}" -- sh
}
