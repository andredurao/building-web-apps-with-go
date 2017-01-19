# Example 11 - Databases

## Databases

On this example we created a simple database with just one table (books) on sqlite3.

The main program take care of create the table, but does not insert nothing on it.

I did that on `sqlite3` CLI utility.

```
sqlite3 example.sqlite
```

And on the database CLI I've did the following insert:
```
sqlite> INSERT INTO books (title, author) VALUES ('Baudolino', 'Umberto Eco');
sqlite> SELECT * FROM books;
Baudolino|Umberto Eco
sqlite>
```

Started a `Dockerfile` to learn how to dockerize apps like this.
