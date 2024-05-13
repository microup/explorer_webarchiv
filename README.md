# explorer_webarchiv
*Download only TEXT data from a specified time period on the web.archive.org*

--------------------------

**When using the downloader, you can specify the following arguments for running it:**

```bash
downloader --domain=YOUR_SITE --timestamp=YOUR_DATE --workers=COUNT_WORKERS
```

where the parameters are:

* domain - specify the target domain (only lowercase)
* timeStamp - specify timestamp in the format:'yyyymmdd' (also: 'yyyy' > download only a specific year; 'yyyymm' > year and month; '2' or '1' > everything for the years past 20** or 19**
* workers - specify the max workers (default=10)

**If you need to build the binary, use the following command:**

```bash
make build
```

----

The original idea for this project was derived from https://github.com/lorenzoromani1983/wayback-keyword-search.
