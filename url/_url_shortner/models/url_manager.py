from url_shortner.errors import ValidationException

from .url import HALF_HOUR, Url
from .user import User

_url_manger = None


class UrlShortnerManager:
    def __init__(self):
        self._users_map = {}
        self._urls_map = {}
        self._anonymous_user = User(name="Anonymous", identity="Anonymous")
        self._users_map[self._anonymous_user.identity] = self._anonymous_user

    def add_user(self, name: str, identity: str):
        if identity in self._users_map:
            raise ValidationException(f"User already exists with identity={identity}")
        user = User(name=name, identity=identity)
        self._users_map[identity] = user

    def add_url_for_user(self, identity, url, ttl, keep_forever) -> Url:
        if ttl is None:
            ttl = HALF_HOUR
        if identity not in self._users_map:
            raise ValidationException(f"User not found for identity={identity}")
        user = self._users_map[identity]
        short_url = self._get_uniq_short_url()
        url = Url(url=url, ttl=ttl, short_url=short_url, keep_forever=keep_forever)
        user.add_url(url)
        self._urls_map[short_url] = {"url": url, "user": user}
        return url

    def add_url_for_anon_user(self, url, ttl) -> Url:
        return self.add_url_for_user(
            identity=self._anonymous_user.identity, url=url, ttl=ttl, keep_forever=False
        )

    def get_url_by_short_url(self, short_url) -> Url:
        if short_url not in self._urls_map:
            raise ValidationException(f"No url found for short_url={short_url}")
        url = self._urls_map[short_url]["url"]
        if url.is_expired:
            self._delete_url(short_url)
            raise ValidationException(f"Url: {short_url} is expired")
        return url

    def _delete_url(self, short_url):
        user = self._urls_map[short_url]["user"]
        user.delete_url(short_url)
        del self._urls_map[short_url]

    def _get_uniq_short_url(self):
        short_url = Url.generate_short_url()
        count = 0
        while short_url in self._urls_map:
            count += 1
            short_url = Url.generate_short_url()
            if count == 10:
                raise ValidationException(f"Unable to build uniq url")
        return short_url


def get_url_manager():
    global _url_manger
    if _url_manger is None:
        _url_manger = UrlShortnerManager()
    return _url_manger
