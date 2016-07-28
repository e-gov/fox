#=============== StaticWeb ===============
FROM node:5.12

COPY ./static /statc/
COPY ./config /conf/

WORKDIR /statc

ARG DEBIAN_FRONTEND=noninteractive

RUN cp /statc/properties.json.sample properties.json \
	&& apt-get update \
	&& apt-get install -y -qq --no-install-recommends rubygems ruby-dev \
	&& gem update --system \
	&& gem install compass --version 1.0.3 \
	&& npm install \
	&& npm install -g grunt-cli

CMD [ "grunt", "serve" ]

