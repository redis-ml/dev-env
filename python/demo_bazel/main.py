import requests


if __name__ == "__main__":
  resp = requests.get("https://google.com")
  print(resp)
  import sys

  print(sys.executable)
