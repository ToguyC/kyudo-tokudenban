Before running the application, you need to create a database and a user in MySQL. The following commands will create a database named `tokudenban` and a user named `MyNiceUser` with the password `MySuperSecurePassword` for example (these credentials must rbe reflected in the `.env` of the application). The user will have all the necessary permissions to run the application.

```sql
CREATE USER 'MyNiceUser'@'%' IDENTIFIED BY 'MySuperSecurePassword';
GRANT Alter ON tokudenban.* TO 'MyNiceUser'@'%';
GRANT Create ON tokudenban.* TO 'MyNiceUser'@'%';
GRANT Create view ON tokudenban.* TO 'MyNiceUser'@'%';
GRANT Delete ON tokudenban.* TO 'MyNiceUser'@'%';
GRANT Delete history ON tokudenban.* TO 'MyNiceUser'@'%';
GRANT Drop ON tokudenban.* TO 'MyNiceUser'@'%';
GRANT Grant option ON tokudenban.* TO 'MyNiceUser'@'%';
GRANT Index ON tokudenban.* TO 'MyNiceUser'@'%';
GRANT Insert ON tokudenban.* TO 'MyNiceUser'@'%';
GRANT References ON tokudenban.* TO 'MyNiceUser'@'%';
GRANT Select ON tokudenban.* TO 'MyNiceUser'@'%';
GRANT Show view ON tokudenban.* TO 'MyNiceUser'@'%';
GRANT Trigger ON tokudenban.* TO 'MyNiceUser'@'%';
GRANT Update ON tokudenban.* TO 'MyNiceUser'@'%';
```

To automate the user creation, please create a file called `user.sql` under the `./docker/mariadb/init/` directory.