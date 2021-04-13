from url_shortner.errors import ValidationException
from url_shortner.services import get_url_shortning_service


class ConsoleClient:
    def __init__(self, url_shortning_service):
        self.url_shortning_service = url_shortning_service

    def run_interactive_mode(self):
        response = None
        while True:
            command_str = input()
            command_and_argument_list = command_str.split(" ")
            command = command_and_argument_list[0]
            argument_list = command_and_argument_list[1:]
            response = self._run_command(command, argument_list)
            if response == "exit":
                break
            print(response)

    def _run_command(self, command, argument_list):
        try:
            if command == "create_short_url_with_ttl":
                assert len(argument_list) >= 1
                url = str(argument_list[0])
                identity = None
                ttl = None
                if len(argument_list) >= 2:
                    ttl = int(argument_list[1])
                if len(argument_list) == 3:
                    identity = str(argument_list[2])

                response = self.url_shortning_service.short_url(
                    url=url, identity=identity, ttl=ttl
                )
            elif command == "create_short_url_without_ttl":
                assert len(argument_list) >= 1
                url = str(argument_list[0])
                identity = None
                if len(argument_list) == 2:
                    identity = str(argument_list[1])
                response = self.url_shortning_service.short_url(
                    url=url, identity=identity, ttl=None
                )
            elif command == "get_url":
                assert len(argument_list) == 1
                short_url = argument_list[0]
                response = self.url_shortning_service.get_url(short_url)
            elif command == "exit":
                response = "exit"
            else:
                response = "Invalid Command"
            return response
        except ValidationException as e:
            return e.args[0]


def get_console_client():
    return ConsoleClient(url_shortning_service=get_url_shortning_service())
