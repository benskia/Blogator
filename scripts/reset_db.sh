#!/bin/bash

echo "----- Resetting database... -----"
goose -dir "./sql/schema" postgres "postgres://postgres:postgres@localhost:5432/blogator" reset
goose -dir "./sql/schema" postgres "postgres://postgres:postgres@localhost:5432/blogator" up

