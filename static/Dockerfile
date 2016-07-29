#=============== StaticWeb ===============
FROM node:5.12

COPY ./static /statc/

WORKDIR /statc

ARG DEBIAN_FRONTEND=noninteractive

RUN cp /statc/properties.json.sample properties.json \
	&& apt-get update \
	&& npm install \
	&& npm install -g grunt-cli

CMD [ "grunt", "serve" ]

