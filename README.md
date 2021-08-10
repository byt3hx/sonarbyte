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

![](https://raw.githubusercontent.com/channyein1337/Sonarbyte/main/image/Sonarbytes.png)

For a single domain
```
▶ echo hackerone.com | sonarbyte 
```
For multiple domains
```
▶ cat domain.txt | sonarbyte
```
Run with httpx
```
▶ cat domain.txt | sonarbyte | sort -u | httpx
```

# Implement

https://omnisint.io/
