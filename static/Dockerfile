#=============== StaticWeb ===============
FROM node:argon

COPY ./static /statc/

RUN cp /statc/properties.json.sample properties.json

WORKDIR /statc

RUN npm install && npm install -g grunt-cli

CMD [ "grunt", "serve" ]

