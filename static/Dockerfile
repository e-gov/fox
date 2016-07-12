#=============== Front End ===============
FROM node:argon

COPY ./static /statc/
COPY ./config /conf/

RUN echo "\n\n\n" && ls -lh /conf && echo "\n\n\n"
RUN echo "\n\n\n" && ls -lh /statc && echo "\n\n\n"
RUN echo "\n\n\n" && ls -lh && echo "\n\n\n"

RUN mkdir -p /conf/Rocker \
	&& cp -n /conf/config.json.template conf/Rocker/config.json \
	&& cp /statc/properties.json.sample properties.json

WORKDIR /statc

RUN npm install && npm install -g grunt-cli

EXPOSE 8090

CMD [ "grunt", "serve" ]

