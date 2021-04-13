from url_shortner.models.url_manager import get_url_manager

_url_shortning_service = None


class UrlShortningService:
    def __init__(self, url_manager):
        self._url_manager = url_manager
        self._create_dummy_users()

    def _create_dummy_users(self):
        dummy_users = {"PersonA": "PersonA", "PersonB": "PersonB"}
        for name, identity in dummy_users.items():
            self._url_manager.add_user(name, identity)

    def short_url(self, identity, url, ttl):
        if identity is None:
            url = self._url_manager.add_url_for_anon_user(url, ttl=None)
            return f"Short url for {url.url} is {url.short_url}"
        keep_forever = True if ttl is None else False
        url = self._url_manager.add_url_for_user(identity, url, ttl, keep_forever)
        return f"Short url for {url.url} is {url.short_url}"

    def get_url(self, short_url):
        url = self._url_manager.get_url_by_short_url(short_url)
        return f"Url for {url.short_url} is {url.url}"


def get_url_shortning_service():
    global _url_shortning_service
    if _url_shortning_service is None:
        _url_shortning_service = UrlShortningService(url_manager=get_url_manager())
    return _url_shortning_service
