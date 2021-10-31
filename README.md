# sonarbyte
[![GitHub issues](https://img.shields.io/github/issues/channyein1337/sonarbyte)](https://github.com/channyein1337/sonarbyte/issues)
[![GitHub forks](https://img.shields.io/github/forks/channyein1337/sonarbyte)](https://github.com/channyein1337/sonarbyte/network)
[![GitHub stars](https://img.shields.io/github/stars/channyein1337/sonarbyte)](https://github.com/channyein1337/sonarbyte/stargazers)

# Description
Sonarbyte is a simple and fast subdomain scanner written in go to extract subdomains from Rapid7's DNS Database using omnisint's api.

# Install
```
▶ go get github.com/channyein1337/sonarbyte
```
```
▶ git clone https://github.com/channyein1337/sonarbyte.git
▶ go build main.go
▶ sudo mv sonarbyte /usr/bin/sonarbyte
```
# Usage
```
▶ sonarbyte -h
Usage of sonarbyte:
  -d string
        Set the domain name
  -df string
        Set the file
```
![](https://raw.githubusercontent.com/channyein1337/Sonarbyte/main/image/Sonarbytes.png)

For a single domain
```
▶ sonarbyte -d hackerone.com 
```
For multiple domains
```
▶ sonarbyte -df domain.txt
```
Run with httpx
```
▶ sonarbyte -df domain.txt | httpx
```

# Updated Feature
- Add new flags


# Implement

https://omnisint.io/
