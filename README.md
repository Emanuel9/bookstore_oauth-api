# bookstore_oauth-api
OAuth API


``` CREATE KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy', 'replication_factor':1};  ```

```
    CREATE TABLE oauth.access_tokens(
       access_token varchar PRIMARY KEY,
       user_id bigint,
       client_id bigint,
       expires bigint     
    );
```