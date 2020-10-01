FROM php:7.4-fpm
RUN apt-get update && apt-get install -y \
        libfreetype6-dev \
        libjpeg62-turbo-dev \
        libpng-dev \
        lighttpd\
        libpq-dev \
        git zip unzip\
    && docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install -j$(nproc) gd \
    && docker-php-ext-configure pgsql -with-pgsql=/usr/local/pgsql \
    && docker-php-ext-install pgsql pdo_pgsql
WORKDIR /var/www/
#ENV TIMEZONE Europe/Moscow
ENV PHP_MEMORY_LIMIT 256M
ENV MAX_UPLOAD 32M
ENV PHP_MAX_FILE_UPLOAD 32
ENV PHP_MAX_POST 32M

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

#COPY ./composer.json /var/www/composer.json
#RUN php /usr/local/bin/composer install
#RUN chown -R www-data:www-data /var/www/
#USER www-data
#EXPOSE 4001
EXPOSE 9000
#CMD ["php-fpm"]
#CMD  php-fpm -D && lighttpd -D -f /etc/lighttpd/lighttpd.conf
COPY ./entrypoint.sh /usr/bin/
RUN ln -s /usr/bin/entrypoint.sh / # backwards compat
RUN chmod a+x /usr/bin/entrypoint.sh
CMD ["/usr/bin/entrypoint.sh"]