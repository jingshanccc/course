# Consul


# MariaDB

`
docker run --name mariadb -v /var/lib/mysql:/var/lib/mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root@123456 -e MYSQL_ROOT_HOST=% -d mariadb --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
`

