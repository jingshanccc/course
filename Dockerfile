FROM alpine
ARG module=user
ADD $module /$module
ENTRYPOINT ["/$module"]