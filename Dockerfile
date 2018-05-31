#######################
# ToRest Docker Image #
#######################

# This image should only be built using the Makefile included
# inside this repository, since that Makefile is in charge of
# preparing all the needed files for the image.

# Debian is only around 80 MB so it is the lightest image we
# can use. There is also the Alpine Linux image weighing only
# 5 MB, but it does not have Glibc, which is required by the
# go-sqlite3 dependency.
FROM debian

# Exposed ports
EXPOSE 8080

# Move caddy, Caddyfile and app files to the image
WORKDIR /srv
COPY files .

# Run the app
ENTRYPOINT ["/srv/caddy", "-conf", "/srv/Caddyfile"]
