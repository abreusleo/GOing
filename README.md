# GOing

Using goroutines and channels for fun

## Querys
(Just so I can ignore migrations for now)

```CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(100) NOT NULL
);```

```INSERT INTO users (name) VALUES (UUID());```

```INSERT INTO users (name)
SELECT UUID()
FROM users;```