# Use the official PHP with Apache image
FROM php:7.4-apache

# Install cURL extension for PHP
RUN apt-get update && \
    apt-get install -y libcurl4-openssl-dev && \
    docker-php-ext-install curl

# Copy the PHP script to the container
COPY api.php /var/www/html/frontend/api.php

# Copy the HTML file to the container
COPY index.html /var/www/html/frontend/index.html

# Expose port 80 for the web server
EXPOSE 80
