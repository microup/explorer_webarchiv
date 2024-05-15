# Explorer web.archiv.org
*The utility allows you to download all textual information for the specified period for the selected URL from web.archive.org.*

--------------------------

**When using the downloader, you can specify the following arguments for running it:**

```bash
explorer_webarchiv --domain=YOUR_SITE --timestamp=YOUR_DATE --workers=COUNT_WORKERS
```

where the parameters are:

* domain - specify the target domain (only lowercase)
* timeStamp - specify timestamp in the format:'yyyymmdd' (also: 'yyyy' > download only a specific year; 'yyyymm' > year and month; '2' or '1' > everything for the years past 20** or 19**
* workers - specify the max workers (default=10)

**If you need to build the binary, use the following command:**

```bash
make build
```

##Features:

* load textual information for a specified period of time for a given domain
* manage the types of data for textual information collection through a config file

----

The original idea for this project was derived from https://github.com/lorenzoromani1983/wayback-keyword-search.

## License

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

####

Copyright (C) 2022-2024 https://microup.ru

