STATICS="../mock-data/static-data"
JSONS="../mock-data/json-data"
if [[ $* == *"mocked"* ]]
then
    export MOCK_STATICS=$STATICS
    export MOCK_JSONS=$JSONS
else
    export STATICS=$STATICS
    export JSONS=$JSONS
    python3 ./scripts/local.py
fi

if [[ $* == *"airmode"* ]]
then

    export $(egrep -v ‘^#’ ../env/.airenv | xargs)
    cd ../src/
    $(go env GOPATH)/bin/air
else
    docker-compose -f ../docker-compose.yml up
fi