while :
do
    redis-cli quit
    if [ $? -eq 0 ]; then
        echo "Server ready now, start to import data ..."
        cat /tmp/seed.txt | redis-cli
        break
    else
        echo "Server not ready, wait then retry..."
        sleep 3
    fi
done