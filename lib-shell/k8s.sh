if [ "$MY_LIB_K8S_DEFINED" = "" ]; then
  export MY_LIB_K8S_DEFINED=1
  source "/usr/local/opt/kube-ps1/share/kube-ps1.sh"
  # export PS1='$(kube_ps1)'$PS1
fi
