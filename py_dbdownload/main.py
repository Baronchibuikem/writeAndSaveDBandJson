import pymongo
import json
import os
from pathlib import Path
from datetime import datetime
from pymongo import MongoClient


def get_mongo_client(uri):
    """Connect to MongoDB and return the client."""
    return MongoClient(uri)


def insert_data_to_mongodb(data, collection):
    """Insert data into a MongoDB collection."""
    if isinstance(data, list):
        if all(isinstance(item, dict) for item in data):
            collection.insert_many(data)
        else:
            raise TypeError("All items in the list must be dictionaries")
    elif isinstance(data, dict):
        collection.insert_one(data)
    else:
        raise TypeError("Data must be a dictionary or a list of dictionaries")


def load_json_file(file_path):
    """Load JSON data from a file."""
    with open(file_path, "r") as file:
        return json.load(file)


def backup_mongodb(uri, db_name, output_dir):
    """Backup MongoDB collections to JSON files."""
    client = get_mongo_client(uri)
    db = client[db_name]

    output_dir.mkdir(parents=True, exist_ok=True)
    collections = db.list_collection_names()

    for collection_name in collections:
        collection = db[collection_name]
        data = list(collection.find({}))

        for item in data:
            if '_id' in item:
                item['_id'] = str(item['_id'])

        file_path = output_dir / f"{collection_name}.json"
        with open(file_path, "w") as file:
            json.dump(data, file, indent=4, default=default_serializer)

    print(f"Backup completed successfully. Data saved in {output_dir}.")


def restore_mongodb(uri, db_name, dir_path):
    """Restore MongoDB collections from JSON files."""
    client = get_mongo_client(uri)
    database = client[db_name]

    for filename in os.listdir(dir_path):
        if filename.endswith(".json"):
            file_path = os.path.join(dir_path, filename)
            data = load_json_file(file_path)

            collection_name = os.path.splitext(filename)[0]
            collection = database[collection_name]

            try:
                insert_data_to_mongodb(data, collection)
                print(f"Data from {filename} successfully inserted into {collection_name} collection")
            except TypeError as e:
                print(f"Error inserting data from file {filename}: {e}")


def default_serializer(obj):
    """JSON serializer for objects not serializable by default json code."""
    if isinstance(obj, datetime):
        return obj.isoformat()
    raise TypeError(f"Type {type(obj)} not serializable")


if __name__ == "__main__":
    MONGO_URI = "mongodb://localhost:27017"
    BACKUP_DB_NAME = "odeserverlocaltestdb"
    RESTORE_DB_NAME = "odebackup"

    # Define the directories
    BACKUP_DIR = Path.home() / "Documents/mongopy/mongodb_backup"
    RESTORE_DIR = Path.home() / "Downloads/mongodb_backup"

    # Backup MongoDB to JSON files
    backup_mongodb(MONGO_URI, BACKUP_DB_NAME, BACKUP_DIR)

    # Restore MongoDB from JSON files
    restore_mongodb(MONGO_URI, RESTORE_DB_NAME, RESTORE_DIR)
