#!/bin/sh

ADDR='localhost:8080'

# ユーザー追加
RESPONSE=$(curl -s -X 'POST' -d '{"user":{"name":"Bob","password":"12345678"}}' "$ADDR/user")
echo $RESPONSE
USER_ID=$(echo $RESPONSE | jq -r '.user.id')
curl -X 'GET' -H "X-User-Id: $USER_ID" "$ADDR/user/$USER_ID"
curl -X 'PUT' -H "X-User-Id: $USER_ID" -d '{"user":{"name":"Alice","password":"12345678"}}' "$ADDR/user/$USER_ID"
curl -X 'DELETE' -H "X-User-Id: $USER_ID" "$ADDR/user/$USER_ID"


# 企業追加
RESPONSE=$(curl -s -X 'POST' -d '{"user":{"name":"Bob","password":"12345678"}}' "$ADDR/user")
echo $RESPONSE
USER_ID=$(echo $RESPONSE | jq -r '.user.id')
curl -X 'POST' -H "X-User-Id: $USER_ID" -d '{"company":{"name":"GREATE COMPANY"}}' "$ADDR/company"
curl -X 'DELETE' -H "X-User-Id: $USER_ID" "$ADDR/user/$USER_ID"
