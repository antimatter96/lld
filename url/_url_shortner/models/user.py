from dataclasses import dataclass, field
from typing import Dict
from uuid import uuid4

from .url import Url


@dataclass
class User:
    name: str
    identity: str
    user_id: str = field(default_factory=lambda: str(uuid4()))
    map_of_urls: Dict[str, Url] = field(default_factory=lambda: {})

    def add_url(self, url: Url):
        assert url.short_url not in self.map_of_urls
        self.map_of_urls[url.short_url] = url

    def delete_url(self, short_url):
        assert short_url in self.map_of_urls
        del self.map_of_urls[short_url]
