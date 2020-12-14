FROM scratch

COPY cmd/cmd .

EXPOSE 8095

ENTRYPOINT [ "./cmd" ]