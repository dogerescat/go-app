FROM mysql:5.7

ENV TZ Asia/Tokyo

# RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && chown -R mysql:root /var/lib/mysql/

COPY my.cnf /etc/mysql/conf.d/my.cnf

COPY sql /docker-entrypoint-initdb.d/

ENV MYSQL_HOST=localhost
ENV MYSQL_PORT=3306
ENV MYSQL_USER=shuto
ENV MYSQL_PASSWORD=shutoaaa
ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=mysql-db

CMD ["mysqld"]

EXPOSE 3306
