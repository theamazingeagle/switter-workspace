FROM php:7.4-fpm
RUN apt-get update && apt-get install -y \
        libfreetype6-dev \
        libjpeg62-turbo-dev \
        libpng-dev \
        lighttpd\
    && docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install -j$(nproc) gd
WORKDIR /var/www/
#ENV TIMEZONE Europe/Moscow
ENV PHP_MEMORY_LIMIT 256M
ENV MAX_UPLOAD 32M
ENV PHP_MAX_FILE_UPLOAD 32
ENV PHP_MAX_POST 32M

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

#COPY ./src /var/www/
#RUN chown -R www-data:www-data /var/www/
#USER www-data
#EXPOSE 4001
EXPOSE 9000
CMD ["php-fpm"]
#CMD  php-fpm -D && lighttpd -D -f /etc/lighttpd/lighttpd.conf