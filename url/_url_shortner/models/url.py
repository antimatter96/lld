import random
import string
from dataclasses import dataclass, field
from datetime import datetime, timedelta
from typing import List

HALF_HOUR = 1800  # 30 mins are 1800 secs
SHORT_URL_SIZE = 8


@dataclass
class Url:
    url: str
    short_url: str
    created_at: datetime = field(default_factory=lambda: datetime.now())
    keep_forever: bool = field(default_factory=lambda: False)
    ttl: int = field(default_factory=lambda: HALF_HOUR)

    @classmethod
    def generate_short_url(cls, short_url_size=SHORT_URL_SIZE) -> str:
        return "".join(
            random.SystemRandom().choice(
                string.ascii_uppercase + string.ascii_lowercase + string.digits
            )
            for _ in range(short_url_size)
        )

    @property
    def is_expired(self):
        if self.keep_forever:
            return False
        if self.created_at + timedelta(0, self.ttl) < datetime.now():
            return True
        return False
