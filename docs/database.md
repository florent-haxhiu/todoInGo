# Database Choices

## Database of Choice

I decided to go with sqlite3 since it was the easiest to set up and very easy to test with

### Schema

I originally had this

```sql
CREATE TABLE Notes (
    id TEXT PRIMARY KEY,
    title TEXT,
    body TEXT,
    userId INTEGER
);

CREATE TABLE User (
    id TEXT PRIMARY KEY,
    username TEXT,
    password TEXT,
    notes BLOB
);
```

But after seeing how I was trying to do both where I wanted to include the `userId` in the `Notes` and also have `notes` in the `User` schema, this would've created some data inconsitencies so rather I removed the `notes BLOB` from the `User` schema and made the `userId` a foreign key so that I can search for `noteId` and `userId`

```sql
CREATE TABLE Users (
    id TEXT PRIMARY KEY,
    username TEXT,
    password TEXT
);

CREATE TABLE Notes (
    id TEXT PRIMARY KEY,
    title TEXT,
    body TEXT,
    userId TEXT,
    FOREIGN KEY (userId) REFERENCES Users(id)
);
```

This is my schema for both Users and Notes

Some of the benefits to this change:
- Easier to query: You can easily find all notes for a user
- Better data integrity: Foreign key constraints ensure data consistency
- More flexible: You can add/remove notes without modifying the user record
- Better performance: No need to serialize/deserialize BLOB data
- Simpler updates: Can update individual notes without touching the user record
