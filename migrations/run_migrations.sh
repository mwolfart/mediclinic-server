#!/bin/bash

goose postgres "host=db port=5432 user=postgres password=postgres dbname=postgres sslmode=disable" up