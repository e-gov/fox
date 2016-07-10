#=============== Front End ===============
FROM node:argon

ADD . /static
WORKDIR /static

RUN mkdir -p /config/Rocker \
	&& apk add --update coreutils \
	&& cp -n /config/config.json.template config/Rocker/config.json

VOLUME /tmp/foxdb

RUN cp config/properties.json.sample properties.json

RUN npm install && npm install -g grunt-cli

EXPOSE 8090

CMD [ "grunt", "serve" ]



