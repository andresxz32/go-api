#!/bin/bash
for i in {0..1000000}
do
  echo $i
  curl -X GET http://localhost:8000/api/v1
done
