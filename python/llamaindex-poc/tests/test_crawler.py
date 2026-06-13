import requests_mock
from ingestion.site_map_crawler import SiteMapCrawler


def test_fetch_urls_standard_sitemap():
    """Test standard sitemap parsing with text extraction, cleaning, and deduplication."""
    sitemap_url = "https://example.com"

    # Mock XML payload containing normal URLs, duplicate layout formats, and asset noise
    mock_xml = """<?xml version="1.0" encoding="UTF-8"?>
    <urlset xmlns="http://sitemaps.org">
        <url><loc>https://example.com</loc></url>
        <url><loc>https://example.com/</loc></url> <!-- Duplicate with trailing slash -->
        <url><loc>https://example.com</loc></url>
        <url><loc>https://example.com</loc></url> <!-- Should be ignored -->
    </urlset>
    """

    with requests_mock.Mocker() as mock:
        mock.get(sitemap_url, text=mock_xml)

        crawler = SiteMapCrawler(sitemap_url=sitemap_url)
        results = crawler.fetch_urls()

        # Assertions
        assert len(results) == 1
        assert results == ["https://example.com"]
