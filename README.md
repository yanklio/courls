## Courls

Small command utility for counting sub-urls of domain by scrapping links on page

# Install

## Unix systems
```
curl -sL https://raw.githubusercontent.com/yanklio/courls/master/install/unix.sh | bash
```

## Windows systems
```
iwr -useb https://raw.githubusercontent.com/yanklio/courls/blob/master/install/windows.ps1 | iex 
```

# Usage

```bash
courls <domain>
```

This will scrape all links from the specified domain and count the unique sub-URLs found. To know additional features please use flag `--help`

# Features

- Fast URL counting and analysis
- Supports domain scraping
- Cross-platform compatibility
- Simple command-line interface

# License

MIT
