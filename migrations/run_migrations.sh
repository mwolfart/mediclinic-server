#!/bin/bash

goose postgres "host=$DB_HOSTNAME port=$DB_PORT user=$DB_USERNAME password=$DB_PASSWORD dbname=$DB_NAME sslmode=require" up