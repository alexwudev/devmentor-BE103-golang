# Installation && Setup
1. Clone the repository and into to directory.
    ```
    git clone git@github.com:jonyig/devmentor-BE103-golang.git && cd devmentor-BE103-golang
    ```

2. Start the container.
   ```sh
   docker-compose up -d 
   ```

3. Check API document
    ```
   localhost:8000
   ```

# Reference

- https://github.com/golang-migrate/migrate

# DB Migrations
## Up
Please run migration up after container mysql is set
```
migration-up
```

## Down
```
migration-down
```

# Example API

## List Post
```
curl --location '127.0.0.1:8080/posts'
```

## Create Post
```
curl --location '127.0.0.1:8080/posts' \
--header 'Content-Type: application/json' \
--data '{
    "title":"12344rr",
    "content" : "1333"
}'
```