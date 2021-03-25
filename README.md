# sqlstrip

When developing you often have to import database dump from e.g. your production servers to debug some problems.
These sql dumps can get quite big and take a long time and much power to import. Depending on your application you can
skip some tables which you don't need for local development (like cache tables).
This tool will accept a list of table names and then strip the `INSERT INTO` lines from the dump before you load it
into MySQL.

## Usage
Pipe an sql dump through sqlstrip and pipe it directly into mysql:
```
cat dump.sql | sqlstrip -table templatecache -table users | mysql dev_table
```

Of course you could also use `> dump_small.sql` to redirect the output into a new file.
