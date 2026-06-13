import os

from dotenv import load_dotenv


# Loading configuration from .env file
def load_config():
    load_dotenv()
    return {"LLAMA_CLOUD_API_KEY": os.getenv("LLAMA_CLOUD_API_KEY")}


CONFIG = load_config()
