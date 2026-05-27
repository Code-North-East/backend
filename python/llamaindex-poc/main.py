from config.settings import CONFIG
from loguru import logger

def main():
    logger.info("Starting server, loading configuration...")
    logger.debug("Using API_KEY: ", CONFIG["LLAMA_CLOUD_API_KEY"])

if __name__ == "__main__":
    main()
