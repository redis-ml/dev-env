IDX_NAME="IdxCreatedAt"

aws dynamodb query \
  --table-name 'post_office_event' \
  --return-consumed-capacity TOTAL \
  --limit '3' \
  --index-name "${IDX_NAME}" \
  --no-scan-index-forward \
  --key-condition-expression "#Key = :v1" \
  --expression-attribute-names '{"#Key": "Owner"}' \
  --expression-attribute-values '{
  ":v1": {"S": "User2"}
}'

echo "##########################"
aws dynamodb query \
  --table-name 'post_office_push_event' \
  --return-consumed-capacity TOTAL \
  --limit '3' \
  --index-name "${IDX_NAME}" \
  --no-scan-index-forward \
  --key-condition-expression "#Key = :v1" \
  --expression-attribute-names '{"#Key": "Owner"}' \
  --expression-attribute-values '{
  ":v1": {"S": "User2"}
}'

echo "##########################"
aws dynamodb query \
  --table-name 'post_office_email_event' \
  --limit '3' \
  --index-name "${IDX_NAME}" \
  --return-consumed-capacity TOTAL \
  --no-scan-index-forward \
  --key-condition-expression "#Key = :v1" \
  --expression-attribute-names '{"#Key": "Owner"}' \
  --expression-attribute-values '{
  ":v1": {"S": "User2"}
}'
