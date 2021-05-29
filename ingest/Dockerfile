# Select Base Image

FROM php:8.0.6-apache

COPY . /var/www/html

RUN pecl install -o -f redis \
    &&  rm -rf /tmp/pear \
    &&  docker-php-ext-enable redis \