# syntax=docker/dockerfile:1
FROM php:8.4.14-apache

WORKDIR /var/www/html

# Copy application source
COPY src/ /var/www/html/
COPY src/users.json /var/www/html/users.json

# Ensure the web server user can read and write the users file
RUN chown www-data:www-data /var/www/html/users.json \
    && chmod 664 /var/www/html/users.json

EXPOSE 80
