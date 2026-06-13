# ingestion/interfaces.py
from abc import abstractmethod, ABC
from typing import List

class BaseSiteMapCrawler(ABC):

    @abstractmethod
    def fetch_urls(self) -> List[str]:
        """
        Discovers, filters and returns a list of strings (targeted urls)
        :return: list of strings
        """
        pass
