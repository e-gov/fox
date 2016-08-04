#=============== StaticWeb ===============
FROM node:5.12

COPY ./static/ /statc/

WORKDIR /statc

RUN cp /statc/properties.json.sample /statc/properties.json \
	&& npm install \
	&& npm install -g grunt-cli

CMD [ "grunt", "serve" ]

