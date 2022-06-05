aws dynamodb get-item \
  --table-name atc \
  --return-consumed-capacity TOTAL \
  --key '{"Owner": {"S": "User2"}, "MsgType": {"S": "v1meta"}}'
