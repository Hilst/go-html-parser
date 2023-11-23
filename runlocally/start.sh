python3 ./scripts/local.py
if [[ $1 == "airmode" ]]
then
        export $(egrep -v ‘^#’ ../env/.airenv | xargs)
        cd ../src/
        $(go env GOPATH)/bin/air
else
    docker-compose -f ../docker-compose.yml up
fi
