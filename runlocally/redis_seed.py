import json
import redis
import os

import util

def make_redis_client() -> redis.client:
    redis_env = util.read_env_file("./redis/.redisenv")
    return redis.Redis(host="localhost",
                       port=6379,
                       db=0,
                       decode_responses=True,
                       password=redis_env["REDIS_PASS"])

def read_json_files() -> dict[str, str]:
   files = {}
   json_folder_path = os.path.join("redis", "seed")
   json_files = [x for x in os.listdir(json_folder_path) if x.endswith("json")]
   for json_file in json_files:
       json_file_path = os.path.join(json_folder_path, json_file)
       with open(json_file_path, "r") as f:
           data = json.load(f)
           files[os.path.splitext(json_file)[0]] = json.dumps(data)
   return files

def seed_data(jsons: dict[str, str], redis_client: redis.client):
   for key, value in jsons.items():
       redis_client.set(key, value)

def redis_main():
    jsons = read_json_files()
    r = make_redis_client()
    seed_data(jsons, r)