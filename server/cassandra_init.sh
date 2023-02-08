#!/usr/bin/env bash

until printf "" 2>>/dev/null >>/dev/tcp/cassandra/9042; do
    sleep 5;
    echo "Waiting for cassandra...";
done

echo "Creating keyspace stocks..."
cqlsh cassandra -e "CREATE KEYSPACE IF NOT EXISTS stocks WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'};"