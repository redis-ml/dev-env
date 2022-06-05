aws dynamodb get-item \
  --table-name po_event \
  --return-consumed-capacity TOTAL \
  --key '{"Owner": {"S": "User2"}, "EventID": {"S": "'"${1}"'"}}'
