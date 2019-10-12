
import sys
import time
import requests as rq
from bs4 import BeautifulSoup

def main():
  f = open("names.txt", "w")
  data = rq.get("https://en.wikipedia.org/wiki/List_of_universities_in_Bangladesh")
  data = BeautifulSoup(data.text, 'html.parser')
  td = data.findAll("td")
  for d in td:
    try:
      a = d.find("a", title=True)
      name = a["title"]
      print(name)
      if name != "STEM fields":
        if name.endswith(" (page does not exist)"):
          name = name[:-len(" (page does not exist)")]
          f.write(f'"{name}",\n')
        else:
          f.write(f'"{name}",\n')
    except Exception as e:
      print(e)
      continue
  f.close()

if __name__ == "__main__":
  main()
