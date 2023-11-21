import subprocess
import time

import minio_screens
import redis_seed

def run_docker_compose(composition: str) -> bool:
    try:
        subprocess.check_call([
            "docker-compose",
            "-f",
            composition + "/docker-compose.yml",
            "-p",
            composition,
            "up",
            "-d"
        ])
        return True
    except:
       return False

def main():
    redis_ok = run_docker_compose("redis")
    minio_ok = run_docker_compose("minio")
    if not redis_ok:
        print("Something wrong with Redis")
        return
    if not minio_ok:
        print("Something wrong with Minio")
        return

    time.sleep(3)

    redis_seed.redis_main()
    minio_screens.minio_main()

main()