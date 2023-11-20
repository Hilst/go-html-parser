import json
from minio import Minio
import os

def make_minio_client() -> Minio:
    return Minio(endpoint="localhost:9000",
                 secure=False,
                 access_key="root",
                 secret_key="password")

def check_buckets(m: Minio):
    if not m.bucket_exists("screens"):
        m.make_bucket("screens")

    folders = [x for x in os.listdir("minio/screens")]
    for folder in folders:
        folder_path = os.path.join("minio/screens", folder)
        for file in os.listdir(folder_path):
            m.fput_object("screens", os.path.join(folder, file), os.path.join(folder_path, file))


def read_only_policy(m: Minio):
    policy = {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": {"AWS": "*"},
                "Action": ["s3:GetBucketLocation", "s3:ListBucket"],
                "Resource": "arn:aws:s3:::screens",
            },
            {
                "Effect": "Allow",
                "Principal": {"AWS": "*"},
                "Action": "s3:GetObject",
                "Resource": "arn:aws:s3:::screens/*",
            },
        ],
    }
    m.set_bucket_policy("screens", json.dumps(policy))


def minio_main():
    m = make_minio_client()
    check_buckets(m)
    read_only_policy(m)