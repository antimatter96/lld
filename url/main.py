from url_shortner.clients import get_console_client


def main():
    console_client = get_console_client()
    console_client.run_interactive_mode()


"""
create_short_url_with_ttl google.com
create_short_url_with_ttl google.com 20 PersonA
create_short_url_with_ttl google.com 10 PersonB
create_short_url_without_ttl google.com PersonB    
get_url 7OUWsW0Q
"""

main()
