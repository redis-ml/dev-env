# Setup Kubernetes with Kop

## Init

1. Setup AWS route53 and S3 with _apply.sh

1. Configure Domain provider by adding corresponding NS records for subdomain.

## Configure

```bash
export KOPS_STATE_STORE=s3://clusters.dev.gods.tools
# Setup configures
kops create cluster --zones=us-west-2b d.dev.gods.tools
```
