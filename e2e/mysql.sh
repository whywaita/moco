#!/bin/bash -e

if [ $# = 0 ]; then
  INSTANCE=0
else
  INSTANCE=$1
fi

if [ -p /dev/stdin ]; then
  MYSQL_OPTS="-u root -p${ROOT_PASSWORD} -e $(cat -)"
else
  MYSQL_OPTS="-u root -p${ROOT_PASSWORD}"
fi

CLUSTER_UID=$(kubectl -n e2e-test get mysqlcluster mysqlcluster -o json | jq -r .metadata.uid)
ROOT_PASSWORD=$(kubectl -n e2e-test get secrets root-password-mysqlcluster-$CLUSTER_UID  -o json | jq -r .data.ROOT_PASSWORD | base64 -d)
kubectl -n e2e-test exec -it mysqlcluster-$CLUSTER_UID-$INSTANCE -- mysql $MYSQL_OPTS
