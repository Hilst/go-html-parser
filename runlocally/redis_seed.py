import json
import redis
import os

def make_redis_client() -> redis.client:
    return redis.Redis(host="localhost",
                       port=6379,
                       db=0,
                       decode_responses=True,
                       password="41b51c8446b802cc442d87ec8e766b6a306afbc35556c2a942ad37fb")

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