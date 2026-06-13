# ingestion/site_map_crawler

import re
from abc import ABC, abstractmethod

import requests
from bs4 import BeautifulSoup
from loguru import logger


class SiteMapCrawler(ABC):
    """
    TODO: add the docstring later...
    """

    def __init__(
            self,
            sitemap_url: str,
            exclude_patterns: list[str] = None,
            timeout: int = 15,
    ):
        """Initializes the sitemap crawler.

        Args:
            sitemap_url (str): The absolute URL to the root sitemap.xml.
            exclude_patterns (List[str], optional): Regex or string fragments
                to ignore (e.g., ['.pdf', '/pricing', '/blog']).
            timeout (int): Network timeout in seconds for fetching XML files.
        """
        self.sitemap_url = sitemap_url
        self.exclude_patterns = exclude_patterns or []
        self.timeout = timeout

        # Standard media/asset extensions to filter out automatically
        self._default_exclusions = [
            r"\.pdf$", r"\.png$", r"\.jpg$", r"\.jpeg$",
            r"\.gif$", r"\.zip$", r"\.xml$", r"#"
        ]
    @abstractmethod
    def fetch_urls(self) -> list[str]:
        """Entry point to extract, filter, and deduplicate documentation URLs.

        Returns:
            List[str]: A unique, list of verified documentation URLs.
        """
        logger.info(f"Starting sitemap discovery at: {self.sitemap_url}")
        discovered_urls: set[str] = set()

        # Recursively parse the root sitemap
        self._parse_sitemap_recursive(self.sitemap_url, discovered_urls)

        # Filter and normalize the discovered batch
        clean_urls = self._filter_and_normalize(discovered_urls)

        logger.info(f"Successfully collected {len(clean_urls)} target URLs.")
        return list(clean_urls)

    def _parse_sitemap_recursive(self, url: str, discovered_urls: set[str]) -> None:
        """Downloads an XML sitemap and checks if it contains sub-sitemaps

        or leaf URLs.
        """
        try:
            response = requests.get(url, timeout=self.timeout)
            response.raise_for_status()
        except requests.RequestException as e:
            logger.error("Failed to fetch sitemap at {}: {}", url, e)
            return

        # Parse XML using BeautifulSoup with the fast 'lxml-xml' parser
        soup = BeautifulSoup(response.content, "lxml-xml")

        # 1. Check if this is a Sitemap Index pointing to sub-sitemaps
        sitemaps = soup.find_all("sitemap")
        if sitemaps:
            for sm in sitemaps:
                loc_tag = sm.find("loc")
                if loc_tag and loc_tag.text:
                    sub_sitemap_url = loc_tag.text.strip()
                    logger.debug("Found nested sub-sitemap: {}", sub_sitemap_url)
                    self._parse_sitemap_recursive(sub_sitemap_url, discovered_urls)
            return

        # 2. Otherwise, treat it as a standard sitemap containing page URLs
        url_tags = soup.find_all("url")
        for ut in url_tags:
            loc_tag = ut.find("loc")
            if loc_tag and loc_tag.text:
                discovered_urls.add(loc_tag.text.strip())

    def _filter_and_normalize(self, raw_urls: set[str]) -> set[str]:
        """Cleans trailing slashes and drops asset noise or restricted paths."""
        processed_urls: set[str] = set()

        # Compile exclusion patterns for performance
        compiled_patterns = [re.compile(p) for p in self.exclude_patterns]
        compiled_defaults = [re.compile(p, re.IGNORECASE) for p in self._default_exclusions]

        for url in raw_urls:
            # Drop empty or malformed structures
            if not url.startswith("http"):
                continue

            # Check against default system asset exclusions (PDFs, PNGs, etc.)
            if any(pattern.search(url) for pattern in compiled_defaults):
                continue

            # Check against user-defined path exclusions (Pricing, Blog, etc.)
            if any(pattern.search(url) for pattern in compiled_patterns):
                continue

            # Normalize trailing slashes to avoid indexing duplicates
            normalized_url = url.rstrip("/")
            processed_urls.add(normalized_url)

        return processed_urls
