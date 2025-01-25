# Database Standards

### 1. Migrations

- **Migrations are planned to be automated by beta.** They are currently executed manually.
- I do not plan on adding an ORM. All migrations will be handwritten.

---

### 2. Naming

- Names shall be `snake_case`.
- All table names should be plural.
- Any table used exclusively for guild purposes will start with the prefix: `guild_`.

---

### 3. Timestamps/Dates

- All stored timestamps and dates will use **Unix time (to the second)**, not milliseconds.
- Refer to the first migration as referance for setting the database default to automatically add the timestamp.

---

### 4. Indexing

- Use an index on any column that is queried for data retrieval.

---

### 5. NOT NULL

- Use NOT NULL when possible.
- "What if it's something unknown till later?", Create a default placeholder that functions so well it doesn't feel like a placeholder.
- Exceptions may occur if it is just simply better to allow null. Example: `category_id` in `guild_channels`; Some channels may be intended to _never_ have a category.

---
