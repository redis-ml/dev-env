# if [ "$MY_LIB_K8S_DEFINED" = "" ]; then
#   export MY_LIB_K8S_DEFINED=1
#   source "$(brew --prefix)/opt/kube-ps1/share/kube-ps1.sh"
#   # export PS1='$(kube_ps1)'$PS1
# fi

function switch_k8s_namespace() {
  local ns=${1?namespace}
  kubectl config \
    set-context \
    $(kubectl config current-context) \
    --namespace="$ns"
}

function switch_k8s_context() {
  local ctxt=${1?context name}
  kubectl config \
    use-context "${ctxt}"
}

function k8s_churn_role() {
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

  local all_pods="$($kubectl_cmd get pods | awk 'substr($1, 1, length(role))==role&&substr($1, length(role)+1)~/^-[0-9A-Za-z]+-[0-9A-Za-z]+$/&&$3=="Running"{printf " pod/"$1;}' role=$role)"
  echo "$kubectl_cmd delete $all_pods"
  $kubectl_cmd delete $all_pods
}

function k8s_search_ip() {
  local ip=${1?ip}
  kubectl get pods --field-selector status.podIP="$ip"  --all-namespaces
}

function k8s_ssh() {
  local role="${1?role}"

  # END of common k8s logic.
  local kubectl_cmd="kubectl"

  local k8s_context_arg=${K8S_CONTEXT:-}
  local k8s_context=""
  if [ "x${k8s_context_arg}" != "x" ]; then
    k8s_context="--context=${k8s_context_arg}"
  fi

  local k8s_namespace_arg=${K8S_NAMESPACE:-}
  local k8s_namespace=""
  if [ "x${k8s_namespace_arg}" != "x" ]; then
    k8s_namespace="--namespace=${k8s_namespace_arg}"
  fi
  # END of common k8s logic.

  # local all_pods="$($kubectl_cmd get pods | awk "\$1~/^$role-[a-f0-9]+-[a-z0-9]+\$/{printf \" pod/\"\$1;}")"
  echo "$(echo $kubectl_cmd $k8s_namespace $k8s_context)"
  local pod_name="$($kubectl_cmd $k8s_namespace $k8s_context get pods | awk 'substr($1, 1, length(role))==role&&substr($1, length(role)+1)~/^-[0-9A-Za-z]+-[0-9A-Za-z]+$/&&$3=="Running"{print $1}' role=$role | head -n 1)"
  echo "pod name: $pod_name"
  $kubectl_cmd $k8s_namespace $k8s_context --container app exec -ti "${pod_name}" -- sh
}
