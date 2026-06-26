import logging


# initiate using the default logging configuration
def init():
    logging.basicConfig(level=logging.DEBUG)

def main():
    init()
    logging.debug("starting opensearch log agent")

if __name__ == "__main__":
    main()