#=============== FrontEnd ===============
FROM node:argon

COPY ./static /statc/
COPY ./config /conf/

RUN mkdir -p /conf/Docker \
	&& cp -n /conf/config.json.template conf/Docker/config.json \
	&& cp /statc/properties.json.sample properties.json

WORKDIR /statc

RUN npm install && npm install -g grunt-cli

CMD [ "grunt", "serve" ]

