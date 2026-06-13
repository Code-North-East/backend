# ingestion/interfaces.py

from abc import ABC, abstractmethod


class BaseSiteMapCrawler(ABC):
    @abstractmethod
    def fetch_urls(self) -> list[str]:
        """
        Discovers, filters and returns a list of strings (targeted urls)
        :return: list of strings
        """
        pass
