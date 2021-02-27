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
